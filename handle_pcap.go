package main

import (
    //"fmt"
    "log"
)





func handlePcap( packet packet_metadata, ipMeta * pState, timeoutQueue * chan packet_metadata, 
    writingQueue * chan packet_metadata, f *output_file ) {

	//verify 
	if !(ipMeta.verifyScanningIP( &packet )) {
		packet.updateTimestamp()
        *timeoutQueue <-packet
		return
	}
    //for every closed connection, record
    if packet.RST || packet.FIN {

        // do not remove packet due to race condition
        //ipMeta.remove(packet)
        *writingQueue <- packet
        //close connection
        if packet.FIN {
            rst := constructRST(packet)
            err = handle.WritePacketData(rst)
        }
        return
     }
     //for every ack received, mark as accepting data
     if (!packet.SYN) && packet.ACK {

         //exit condition
         if len(packet.Data) > 0 {

             //remove from state, we are done now
             ipMeta.remove(packet)
             packet.fingerprintData()
             *writingQueue <- packet
             //close connection
             rst := constructRST(packet)
             err = handle.WritePacketData(rst)
             if err != nil {
                log.Fatal(err)
             }
             return

         }
		 //add to map
		 packet.updateResponse(DATA)
		 packet.updateTimestamp()
		 ipMeta.update(packet)

		 //add to map
         *timeoutQueue <-packet

    }

}

