const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        listOfBalancers: () => client.get('/balancers'),
        updateMachine: (id, worked) => client.post('/updateMachine', { id, worked })
    }

};

module.exports = { Client };
