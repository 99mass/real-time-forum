import { linkApi } from "../helper/api_link.js";


const getComments = (_comments, IdPost,createCommentForm, callback) => {
    for (let i = 0; i < _comments.length; i++) {
        const comment = _comments[i];
        let postId = IdPost[i].textContent;
        comment.addEventListener("click", function () {        
            fetch(`${linkApi}post`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ PostID: postId })
            })
                .then(response => {
                    if (response.status !== 200) {
                      console.error('Error:' + response["message"]);
                      return;
                    }
                    return response.json()
                })
                .then(data => {
                    callback(data);
                })
                .catch((error) => {
                    console.error('Error: '+error);
                });
        });
    }
}

export { getComments }