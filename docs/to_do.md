# 

1) need to put constraints over the leave span applications.

2) auth is shitty. once you get the auth token and session id, you can basically successfully verify million years later or whenever.

3) the username should be unique in mongodb


# basic structure APIs

1)  a) need to show everyone who's on leave on any  day.
    b) count of everyone who's on leave today.
    c) above functionalities should be available for future dates as well, only approved leaves.
    d) same functionalitites should be availabe according to the team.

2)  a) approver should be able to see leave data for each employee under him(reporter)?

    b) For each day which reporter is on leave.

    c) admin should be able to see anyone should be on leave in the whole Org.

3) a) Try to implement cache for reducing  redundant data fetching from database, everytime a find query is requested?

4)  a) Creating a new collection for  public holdiays!!!

    b) Create an API endpoint for it as well.

    c) create only one collection for both public holdydas and weekends. Every holiday would have a special field names `isPublicHoloday`, which would be a bool value, and it would be set to tru,e if it actually is a public holiday. SO, every weekend, that's not a public holdiay would have the `is PublicHoliday` set to `false`. ```YAYYYYYY!!!!!!!!!```

10) ## FRONTEND & BACKEND RESTRICTIONS!!
    a) already defined public holidays should be removed from the leave data time span, and shoul be split into list consisting of the broken leave spans(the list doesn't containt the public holidays.)

    b) max leaves limit should be defined per year. the leaves left unused from pervious year should be added for the next one.

    c) leaves balance should be always equal to or more than the leave applied.

    d) the leaves should be approved before the starting date fo the `leaveData`.

    e) User should be able to apply leaves in a descrete format 

    f) 