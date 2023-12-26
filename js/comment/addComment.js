import { linkApi } from "../helper/api_link.js";
import { commentTemporel } from "../helper/utils.js";
import { liskeComment } from "../likeDislike/comment/like.js";
import { DisLiskeComment } from "../likeDislike/comment/dislike.js";


const addComment=(formAddComment,blocComment)=>{

  for (let i = 0; i < formAddComment.length; i++) {

        const _form = formAddComment[i];    
        _form.addEventListener('submit', async function(e) {
            e.preventDefault();
            let postId = document.querySelectorAll('input[name="postId"]')[i].value.trim();
            let userId = document.querySelectorAll('input[name="userId"]')[i].value.trim();
            let content = document.querySelectorAll('textarea[name="ContentComment"]')[i].value.trim();
            let userName = document.querySelectorAll('input[name="userName"]')[i].value.trim();

            let data = {
                PostID: postId,
                UserID: userId,
                Content: content 
            };
            try {
                const response = await fetch(`${linkApi}addcomment`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(data),
                });
            
                if(response.status === 200){    
                    const data = await response.json();
                    // afficher le derniere commentaire
                    if (data["Comment"][0]) {
                                           
                        let contentCom=data["Comment"][0]["Comment"]["Content"];
                        let dateCom=data["Comment"][0]["Comment"]["CreatedAt"];
                        let idCom=data["Comment"][0]["Comment"]["ID"];
                        let userCom=data["Comment"][0]["User"]["Username"];
                        let newComment=commentTemporel(contentCom,userCom,0,0,idCom,dateCom);
                        
                        blocComment[i].insertAdjacentHTML('afterbegin', newComment);
                        
                        setTimeout(() => {                        
                            const likeComment = document.querySelectorAll('.like-comment-block .like-comment');
                            const dislikeComment = document.querySelectorAll('.like-comment-block .dislike-comment');
                            const likeCommentId = document.querySelectorAll('.id-comment-like');
                            const dislikeCommentId = document.querySelectorAll('.id-comment-dislike');
                            const likeCommentScore = document.querySelectorAll('.like-comment-block .like-comment .scorecommentLike');
                            const dislikeCommentScore = document.querySelectorAll('.like-comment-block .dislike-comment .scorecommentDisLike');
                            
                            liskeComment(likeComment, likeCommentId, likeCommentScore, dislikeCommentScore, dislikeComment);
                            DisLiskeComment(dislikeComment, dislikeCommentId, likeCommentScore, dislikeCommentScore, likeComment);        
                        }, 1000);
                    }

                }else{
                    const data = await response.json();
                    alert('Error : ' + data.message);
                    return;
                }
                
            } catch (error) {
               console.error("error : "+error);
            }

            
            
        });
    } 
}

export {addComment}