#!/bin/bash

CNT=0
ITER=$1
SLEEP=$2
NUMBLOCKS=$3
NODEADDR=$4

if [ -z "$1" ]; then
	echo "Need to input number of iterations to run..."
	exit 1
fi

if [ -z "$2" ]; then
	echo "Need to input number of seconds to sleep between iterations"
	exit 1
fi

if [ -z "$3" ]; then
	echo "Need to input block height to declare completion..."
	exit 1
fi

if [ -z "$4" ]; then
	echo "Need to input node address to poll..."
	exit 1
fi

docker_containers=($(docker ps -q -f name=bitoroprotocold --format='{{.Names}}'))

while [ ${CNT} -lt "$ITER" ]; do
	curr_block=$(curl -s "$NODEADDR":26657/status | jq -r '.result.sync_info.latest_block_height')

	if [ -n "${curr_block}" ]; then
		echo "Number of Blocks: ${curr_block}"
	fi

	if [ -n "${curr_block}" ] && [ "${curr_block}" -gt "${NUMBLOCKS}" ]; then
		echo "Number of blocks reached. Success!"
		exit 0
	fi

	# Check logs for "CONSENSUS FAILURE" on all containers
	for container in "${docker_containers[@]}"; do
		logs=$(docker logs "$container")
		if [[ "$logs" == *"CONSENSUS FAILURE"* ]]; then
			log_line=$(echo "$logs" | grep "CONSENSUS FAILURE")
			echo "Error: Found 'CONSENSUS FAILURE' in logs of container $container: $log_line"
			exit 1
		fi
	done

	# Emulate network chaos:
	# Every $SLEEP * 5 seconds, pick a random container and restart it.
	if ! ((CNT % 5)); then
		rand_container=${docker_containers["$((RANDOM % ${#docker_containers[@]}))"]}
		echo "Restarting random docker container ${rand_container}"
		docker restart "${rand_container}" &>/dev/null &
	fi
	((CNT = CNT + 1))
	sleep "$SLEEP"
done
echo "Timeout reached. Failure!"
exit 1
