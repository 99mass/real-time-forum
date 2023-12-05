import { linkApi } from "../helper/api_link.js";


const  getComments=(_comments,IdPost)=>{
    // let postId
        for (let i = 0; i < _comments.length; i++) {
            const comment = _comments[i];
            let  postId=IdPost[i].textContent;
        comment.addEventListener("click", function () {
        //     try {
        //         const response =  fetch(`${linkApi}post`, {
        //             method: 'POST',
        //             headers: {
        //                 'Content-Type': 'application/json',
        //             },
        //             body: JSON.stringify({PostID: postId}),
                    
        //         });
                
        //         if(response.status === 200){
        //         console.log("ok : "+IdPost[i].textContent);
        //         // window.location.reload();
        //     }
            
        // } catch (error) {
        //    console.log("error : "+error);
        // }
        fetch(`${linkApi}post`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({PostID: postId})
        })
        .then(response =>{ 
            if (response.status===200) {
                // window.location.reload();
console.log('ok');
            }
            return  response.json()
        })
        .then(data =>{ 
            console.log(data)
        })
        .catch((error) => {
            console.error('Error:', error);
        });
    
        });
    }

}

export {getComments}