import {
  generateUsernameForSubaccount,
} from '../../src/helpers/usernames-helper';

const addresses = [
  'btoro1gf4xlnpulkyex74asxxhg9ye05r28cxdd69s9u',
  'btoro10fx7sy6ywd5senxae9dwytf8jxek3t2gcen2vs',
  'btoro1wau5mja7j7zdavtfq9lu7ejef05hm6ffenlcsn',
  'btoro1df84hz7y0dd3mrqcv3vrhw9wdttelul8edqmvp',
  'btoro16h7p7f4dysrgtzptxx2gtpt5d8t834g9dj830z',
  'btoro15u9tppy5e2pdndvlrvafxqhuurj9mnpdstzj6z',
  'btoro1q54yvrslnu0xp4drpde6f4e0k2ap9efss5hpsd',
  'btoro1n88uc38xhjgxzw9nwre4ep2c8ga4fjxc565lnf',
  'btoro1n88uc38xhjgxzw9nwre4ep2c8ga4fjxc575lnf',
  'btoro199tqg4wdlnu4qjlxchpd7seg454937hjrknju4',
  'btoro1c05rgh22wg5pmaufnj8z77kla8fgrgkrth6hlj',
  'btoro1qvsqtewdxqdpj5gnkl8usqz3yplpkgs52wv4fy',
];

describe('usernames-helper', () => {
  it('Check result of generated username', () => {
    const expectedUsernames = [
      'CurlyHallVE6',
      'AmpleCrownLKI',
      'LoneLawXWV',
      'BlueGapOGY',
      'ZippyElfQE0',
      'WindyFaxLEC',
      'RoyalFruit59E',
      'GreenSnowWTT',
      'BubblyEarH5Y',
      'LunarMatFK5',
      'SubtleDig25M',
      'HillyAccess1C7',
    ];

    const gotUserNames = [];
    for (let i = 0; i < addresses.length; i++) {
      const address = addresses[i];
      const namesForOneAddress = new Set();
      for (let k = 0; k < 10; k++) {
        const username: string = generateUsernameForSubaccount(address, 0, k);
        if (k === 0) {
          gotUserNames.push(username);
        }
        namesForOneAddress.add(username);
      }
      // for same address, difference nonce should result in different username
      expect(namesForOneAddress.size).toEqual(10);
    }
    expect(gotUserNames).toEqual(expectedUsernames);
  });

  it('Check determinism of generated username', () => {
    for (let i = 0; i < addresses.length; i++) {
      const address = addresses[i];
      const namesForOneAddress = new Set();
      for (let k = 0; k < 10; k++) {
        const username: string = generateUsernameForSubaccount(address, 0, 0);
        namesForOneAddress.add(username);
      }
      // for same address, difference nonce should result in different username
      expect(namesForOneAddress.size).toEqual(1);
    }
  });
});
