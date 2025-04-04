import {
  SubaccountUsernamesTable,
  SubaccountTable,
  QueryableField,

  testMocks,
  dbHelpers,
  SubaccountUsernamesFromDatabase,
  SubaccountFromDatabase,
} from '@bitoroprotocol-indexer/postgres';
import subaccountUsernameGenerator from '../../src/tasks/subaccount-username-generator';

describe('subaccount-username-generator', () => {
  beforeAll(async () => {
    await dbHelpers.migrate();
    await dbHelpers.clearData();
  });

  beforeEach(async () => {
    await testMocks.seedData();
    await testMocks.seedAdditionalSubaccounts();
  });

  afterAll(async () => {
    await dbHelpers.teardown();
    jest.resetAllMocks();
  });

  afterEach(async () => {
    await dbHelpers.clearData();
    jest.clearAllMocks();
  });

  it('Successfully creates a username for all subaccount', async () => {
    const subaccounts: SubaccountFromDatabase[] = await
    SubaccountTable.findAll({
      subaccountNumber: 0,
    }, [QueryableField.SUBACCOUNT_NUMBER], {});

    const subaccountsLength: number = subaccounts.length;
    const subaccountsWithUsernames: SubaccountUsernamesFromDatabase[] = await
    SubaccountUsernamesTable.findAll(
      {}, [], {});
    expect(subaccountsWithUsernames.length).toEqual(0);

    await subaccountUsernameGenerator();
    const subaccountsWithUsernamesAfter: SubaccountUsernamesFromDatabase[] = await
    SubaccountUsernamesTable.findAll(
      {}, [], {});

    const expectedUsernames = [
      'BubblyEarH5Y', // btoro1n88uc38xhjgxzw9nwre4ep2c8ga4fjxc575lnf
      'GreenSnowWTT', // btoro1n88uc38xhjgxzw9nwre4ep2c8ga4fjxc565lnf
      'LunarMatFK5', // btoro199tqg4wdlnu4qjlxchpd7seg454937hjrknju4
    ];
    expect(subaccountsWithUsernamesAfter.length).toEqual(subaccountsLength);
    for (let i = 0; i < expectedUsernames.length; i++) {
      expect(subaccountsWithUsernamesAfter[i].username).toEqual(expectedUsernames[i]);
    }
  });
});
