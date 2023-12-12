import { linkApi } from "../helper/api_link.js";

let isThrottled = false;

const userOnline = () => {
    // var containUserOnLine=document.querySelector('.bloc-users-on-line');
    fetch(`${linkApi}connectedUsers`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
    })
    .then(response => {
        if (response.status === 200 ) {                       
            return response.json();
        }
    })
    .then(data => {
        console.log(data);
    })
    .catch(error => {
        console.error('Error:', error);
    });
}

export { userOnline }