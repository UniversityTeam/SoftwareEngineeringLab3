// This file contains examples of scenarios implementation using
// the SDK for channels management.

const channels = require('./channels/client');

const client = channels.Client('http://localhost:8080');

// Scenario 1: Display available channels.
client.listOfBalancers()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Available balancers:');
        console.log(list);
    })
    .catch((e) => {
        console.log(`Problem listing available balancers: ${e.message}`);
    });

// Scenario 2: Create new channel.
client.updateMachine(3, true)
    .then((resp) => {
        console.log('=== Scenario 2 ===');
        console.log('Update machine status response:', resp);
    })
    .catch((e) => {
        console.log(`Problem creating a new channel: ${e.message}`);
    });
