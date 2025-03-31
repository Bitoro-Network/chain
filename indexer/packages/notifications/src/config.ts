/**
 * Environment variables required for Notifications module.
 */

import {
  parseString,
  parseSchema,
  baseConfigSchema,
} from '@bitoroprotocol-indexer/base';

export const notificationsConfigSchema = {
  ...baseConfigSchema,

  // Private Key for the Google Firebase Messaging project
  // default is a dummy value
  FIREBASE_PRIVATE_KEY_BASE64: parseString({ default: 'LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tClBMQUNFSE9MREVSX0tFWV9GT1JfREVWRUxPUE1FTlQKLS0tLS1FTkQgUFJJVkFURSBLRVktLS0tLQo=' }),

  // APP ID for the Google Firebase Messaging project
  // default is a dummy value
  FIREBASE_PROJECT_ID: parseString({ default: 'bitoro' }),

  // Client email for the Google Firebase Messaging project
  // default is a dummy value
  FIREBASE_CLIENT_EMAIL: parseString({ default: 'firebase-adminsdk-f0joo@bitoro.iam.gserviceaccount.com' }),
};

export default parseSchema(notificationsConfigSchema);
