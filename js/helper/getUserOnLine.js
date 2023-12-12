import { linkApi } from "../helper/api_link.js";

let isThrottled = false;

const userOnline = async () => {

        const response = await fetch(`${linkApi}connectedUsers`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        if (response.status === 200 && response.headers.get('content-type').includes('application/json')) {                       
            // console.log( await response.json());
            return await response.json();
            
        }

}



export { userOnline }