The tech team, going forward, would like to be able to keep track of every standup update. As such, daily standups would instead of being verbal updates, be written updates that would be recorded within the 15-minute window that is supposed to be used for standups. In other words, instead of waking up to join a call, you'll wake up to provide your update within the 15mins window allocated for standups.

Build a simple API-enabled application (purely backend ) with Golang, that will;
● Allow employees to provide their daily updates
● Allow employees to see all the updates for every one
● Allow employees to filter updates by week/day/sprint/owner (who wrote it )
● A single check-in/update will include the following data
    ● EmployeeID
    ● Employee name
    ● Date
    ● Sprint ID
    ● Task IDs ( can be a text )
    ● What was done yesterday
    ● What will be done today
    ● Blocked by employee IDs - which employees are blocking this person
    ● Breakaway
    ● Check-in time
    ● Status (before standup, after standup, within standup ) around what time was this logged.
NOTE: No deletion or updating of updates is allowed.