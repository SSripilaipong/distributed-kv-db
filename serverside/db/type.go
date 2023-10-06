package db

type Func func(dbPort, peeringPort int, advertisedIp string, peerAddresses []string) error
