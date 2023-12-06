import { linkApi } from "../helper/api_link.js";

const addComment=(formAddComment)=>{
  for (let i = 0; i < formAddComment.length; i++) {

        const _form = formAddComment[i];    
        _form.addEventListener('submit', async function(e) {
            e.preventDefault();
            let postId = document.querySelectorAll('input[name="postId"]')[i].value.trim();
            let userId = document.querySelectorAll('input[name="userId"]')[i].value.trim();
            let content = document.querySelectorAll('textarea[name="ContentComment"]')[i].value.trim();
            if (content==="") {
                alert('The message is required');
            }
            let data = {
                PostID: postId,
                UserID: userId,
                Content: content 
            };
            console.log(data);
            try {
                const response = await fetch(`${linkApi}addcomment`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(data),
    
                });
    
                if(response.status === 200){                   
                    window.location.reload();
                }
                
            } catch (error) {
               console.log("error : "+error);
            }


        });
    } 
}

export {addComment}