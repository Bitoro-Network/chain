import { logger, stats } from '@bitoroprotocol-indexer/base';
import {
  createNotification,
  NotificationDynamicFieldKey,
  NotificationType,
  sendFirebaseMessage,
} from '@bitoroprotocol-indexer/notifications';
import {
  OrderFromDatabase,
  PerpetualMarketFromDatabase,
  SubaccountFromDatabase,
  SubaccountTable,
  FirebaseNotificationTokenTable,
} from '@bitoroprotocol-indexer/postgres';

import config from '../../config';

export async function sendOrderFilledNotification(
  order: OrderFromDatabase,
  market: PerpetualMarketFromDatabase,
) {
  const start = Date.now();
  try {
    const subaccount = await SubaccountTable.findById(order.subaccountId);
    if (!subaccount) {
      throw new Error(`Subaccount not found for id ${order.subaccountId}`);
    }

    const tokens = (await FirebaseNotificationTokenTable.findAll(
      { address: subaccount.address }, [])
    );
    if (tokens.length === 0) {
      return;
    }

    const notification = createNotification(
      NotificationType.ORDER_FILLED,
      {
        [NotificationDynamicFieldKey.AMOUNT]: order.size.toString(),
        [NotificationDynamicFieldKey.MARKET]: market.ticker,
        [NotificationDynamicFieldKey.AVERAGE_PRICE]: order.price,
      },
    );

    await sendFirebaseMessage(tokens, notification);
  } catch (error) {
    logger.error({
      at: 'ender#notification-functions',
      message: 'Error sending order filled notification',
      error,
    });
  } finally {
    stats.timing(`${config.SERVICE_NAME}.send_order_filled_notification.timing`, Date.now() - start);
  }
}

export async function sendOrderTriggeredNotification(
  order: OrderFromDatabase,
  market: PerpetualMarketFromDatabase,
  subaccount: SubaccountFromDatabase,
) {
  const start = Date.now();
  try {
    const tokens = (await FirebaseNotificationTokenTable.findAll(
      { address: subaccount.address }, [],
    ));
    if (tokens.length === 0) {
      return;
    }
    const notification = createNotification(
      NotificationType.ORDER_TRIGGERED,
      {
        [NotificationDynamicFieldKey.MARKET]: market.ticker,
        [NotificationDynamicFieldKey.PRICE]: order.price,
        [NotificationDynamicFieldKey.AMOUNT]: order.size.toString(),
      },
    );
    await sendFirebaseMessage(tokens, notification);
  } catch (error) {
    logger.error({
      at: 'ender#notification-functions',
      message: 'Error sending order triggered notification',
      error,
    });
  } finally {
    stats.timing(`${config.SERVICE_NAME}.send_order_triggered_notification.timing`, Date.now() - start);
  }
}
