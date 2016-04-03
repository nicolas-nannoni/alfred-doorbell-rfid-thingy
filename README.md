# Alfred Doorbell RFID - Fingy device implementation
Alfred is an access control service, that can manage a set of users, sites, doors, locks and more. Administrators can set which users can access which sites, which areas, which doors, potentially affected by schedules and other events.

This package is the device implementation made to work with a doorbell system. It is designed to run on a Raspberry-like device, connected to an Arduino-like device. It connects to the Alfred remote service via Fingy, and receives information about swiped badges, entered PIN and doorbell notification via the serial connection established with the Arduino.

**NOTE: this is still a pre-alpha version. Everything can change and documentation is close to non-existing.**

## How to use?
For now, just run the executable. You might have to change some parameters inside (`deviceId`, `fingyGatewayHostname`), this will change later.

## Remaining work
* Package (configuration, documentation) this nicely...
* Connect to Alfred via Fingy and handle swiped card, PIN and doorbell events.
# Offline/Backup mode (master password/card that can unlock the door without remote Alfred access)
* Assisted registration UI locally on the device, to securely register the device in Fingy without having to know its IP (connection to a phone?)