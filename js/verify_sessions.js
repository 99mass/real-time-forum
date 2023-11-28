const isSessionFound = () => {
    // let sessionId = sessionStorage.getItem('sessionID');
    let sessionId = getCookie('sessionID');
    if (sessionId) {
        console.log("Session found: " + sessionId);
        fetch('http://localhost:8080/verifySession', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({session: sessionId}),
        })
        .then(response => {
            
            if (response.headers.get('content-type').includes('application/json')) {
                return response.json();
            } else {
                console.log(response.status);
                throw new Error('Received non-JSON response');
            }
        })
        .then(data => console.log(data))
        .catch((error) => {
          console.error('Error:', error);
        });
    } else {
        console.log("No session found");
    }
}
function getCookie(name) {
    let cookieArr = document.cookie.split(";");
    for(let i = 0; i < cookieArr.length; i++) {
        let cookiePair = cookieArr[i].split("=");
        if(name == cookiePair[0].trim()) {
            return decodeURIComponent(cookiePair[1]);
        }
    }
    return null;
}

export {isSessionFound}