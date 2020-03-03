package main

import (
    "log"
    "fmt"
	"encoding/json"
)



func ackZMap(input string, ipMeta * pState, timeoutQueue * chan packet_metadata ) {

        fmt.Println(input)

        var synack packet_metadata
        //expecting ip,sequence number, acknumber,windowsize
        err = json.Unmarshal( []byte(input),&synack )
        if err != nil {
            log.Fatal(err)
        }

        //TODO: check that ip_metadata contains what we want (saddr,seq,ack,window)

        if synack.windowZero() {
            //not a real s/a
            return
        }

        //Send Ack with Data
        ack := constructAck(synack)
        fmt.Println("Constructed ack...")

		//add to map
		synack.updateState(ACK)
		synack.updateTimestamp()
		ipMeta.update(synack)

        err = handle.WritePacketData(ack)
        if err != nil {
            log.Fatal(err)
        }
        *timeoutQueue <-synack
		return

}


