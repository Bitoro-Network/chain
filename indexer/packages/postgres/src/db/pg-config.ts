import moment from 'moment';
import pg from 'pg';

/**
 * we need to add this line, because the default type parser
 * changes all datetime objects to javascript dates, when we
 * need all dates returned with iso strings
 */
pg.types.setTypeParser(
  pg.types.builtins.TIMESTAMPTZ,
  (val) => (val === null ? null : moment(val).toISOString()),
);
