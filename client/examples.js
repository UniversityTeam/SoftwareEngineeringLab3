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

setTimeout(() => {
    // Scenario 2: Update a machine.
    client.updateMachine(2, false)
        .then((resp) => {
            console.log('=== Scenario 2 ===');
            console.log('Update machine status response:', resp);
        })
        .catch((e) => {
            console.log(`Problem updating a machine: ${e.message}`);
        });
}, 2000)

setTimeout(() => {
    // Scenario 3: Display available balancers.
    client.listOfBalancers()
        .then((list) => {
            console.log('=== Scenario 1 ===');
            console.log('Available balancers after update machine:');
            console.log(list);
        })
        .catch((e) => {
            console.log(`Problem listing available balancers: ${e.message}`);
        });
}, 3000)