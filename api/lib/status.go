package lib

/** users status */
const USER_PENDING = 0     // never do payment success
const USER_ACTIVE = 1      // paid, payment success
const USER_VERIFIED = 2    // verified account alias email
const USER_COMPLETED = 3   // completed user profile
const USER_EARNING = 4     // do withdraw
const USER_SUSPEND = 5     // suspend account
const USER_TOBEDELETED = 6 // to be deleted account according time setting

/** users category */
const CATEGORY_SUPER_ADMIN = 1
const CATEGORY_ADMIN = 2
const CATEGORY_USER = 3
