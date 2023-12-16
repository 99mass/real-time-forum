import { sendMessages, recipientMessages } from "../layout/corps.js";
import { chatDateFormatter } from "../helper/utils.js";


const OldChatMessage = (_User1, _User2, chatBody) => {
    console.log('user 1: ' + _User1);
    console.log('user 2: ' + _User2);
    console.log(chatBody);

    const socket = new WebSocket("ws://localhost:8080/communication");

    let Messdata = {
        User1: _User1,
        User2: _User2,
    }
    // Send user connected message when connection is open

    socket.onopen = () => {
        socket.send(JSON.stringify(Messdata));
        console.log("WebSocket communication on.");
    }

    socket.onmessage = (message) => {
        var _datas = JSON.parse(message.data);
        console.log(_datas);
        chatBody.innerHTML = "";

        if (_datas) {
            _datas.forEach(_data => {

                let formattedDate = chatDateFormatter(_data["Created"]);

                if (_data["Sender"] == _User1) {
                    let send = sendMessages(_data["Sender"], _data["Message"], formattedDate);
                    chatBody.appendChild(send);

                }

                if (_data["Recipient"] == _User2) {
                    let recip = recipientMessages(_data["Sender"], _data["Message"], formattedDate);
                    chatBody.appendChild(recip);
                }
            });
        }
    };
}

export { OldChatMessage }