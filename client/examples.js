// This file contains examples of scenarios implementation using
// the SDK for balancers management.

const balancers = require('./balancers/client');

const client = balancers.Client('http://localhost:8080');

// Scenario 1: Display available balancers.
client.listOfBalancers()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Available balancers:');
        console.log(list);
    })
    .catch((e) => {
        console.log(`Problem listing available balancers: ${e.message}`);
    });

// Scenario 2: Update a machine.
client.updateMachine(3, true)
    .then((resp) => {
        console.log('=== Scenario 2 ===');
        console.log('Update machine status response:', resp);
    })
    .catch((e) => {
        console.log(`Problem updating a machine: ${e.message}`);
    });
