module.exports = {
    networks: {
        development: {
            host: "127.0.0.1",
            port: 7545,
            network_id: "*",
            gas: 4700000
        },
        testnet: {
            host: "127.0.0.1",
            port: 8545,
            network_id: "*",
            gas: 47000000
        }
    },
    solc: {
        optimizer: {
            enabled: true,
            runs: 200
        }
    }
}
