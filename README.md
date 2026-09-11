# Tide Tracker


## Description
It emails the recipients if the tide reaches a certain height


## Getting Started

This requires some env variables to use:

* USER - user sending the emails - plain string email address
* PASSWORD - application password to allow the smtp email to work
* RECIPIENTS - Space separated email addresses to notify
* LOCATION - the location to notify the tide for

example:
```
USER: my.email@gmail.com
PASSWORD: mypassword
RECIPIENTS: person.1@gmail.com person.2@hotmail.com person3@yahoo.com
LOCATION: my-location-tide-times
```

### Executing program

```
go run .
```