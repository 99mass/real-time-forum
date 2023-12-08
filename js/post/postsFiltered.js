
const postsFilter=(_posts) => {
    var posts="",lastPost="",lastFormComment="",lastBlocComment="";
    for (let i = 0; i < _posts.length; i++) {
        const post = _posts[i];
        const title=post["Posts"]["Title"];
        const content=post["Posts"]["Content"];
        const Categories=post["Posts"]["Categories"];
        const Username=post["User"]["Username"];
        const Image=post["Posts"]["Image"];
        const commentCount=post["CommentCount"];
        const postLikeCount=post["PostLike"];
        const postDislikCount=post["PostDislike"];
        const userId=post["Posts"]["UserID"];
        const postId=post["Posts"]["ID"];
        let specifiqueComment=post["Comment"] ? post["Comment"] : "";
       

        if (i===_posts.length-1 ){
            lastPost="last-post";
            lastFormComment="last-form-comment";
            lastBlocComment="last-bloc-comment";
        }

        // recupere les categorie du post
        let NameCategories="";
        for (let j = 0; j < Categories.length; j++) {
            const NameCategory = Categories[j]["NameCategory"];
            NameCategories+=`<span>${NameCategory}</span>`;
        }
        
    let image=Image ? `<div class="image-post"><img src="/static/image/${Image}" alt=""></div>` : "";
    let _comme="";
    
    // remplire les commentaires
    if (specifiqueComment!=="") {   
        for (let k = 0;k < specifiqueComment.length; k++) {
            const spComment = specifiqueComment[k];  
            
            const ComContent= spComment["Comment"]["Content"];
            const ComCreatedAt= spComment["Comment"]["CreatedAt"];
            const ComPostId= spComment["Comment"]["PostID"];
            const ComUserId= spComment["Comment"]["UserID"];
            const ComLike=spComment["CommentLike"];
            const ComDislike=spComment["CommentDislike"];
            const ComUserName=spComment["User"]["Username"];             
            _comme+=`<div >
                        <div class="comment-text "><pre class="card-description">${ComContent}</pre></div>
                        <div class="content-comment-like">
                        <div class="content-comment">
                            <div class="comment">                                  
                                <img src="assets/user-profile-svgrepo-com.svg" alt="">
                                <div>
                                    <p><span>${ComUserName}</span> </p>
                                    <p>2 weeks ago</p>
                                </div>
                            </div>
                            <div class="like-comment-block">
                                <div>${ComDislike} dislikes</div>
                                <div>${ComLike} Likes</div>
                            </div>
                        </div>
                    </div>
                        </div>
                    <hr>
            `;
        } 
     } 


     posts += `<div class="one-post-block ${lastPost}">
                ${image}
                <div class="post-content">
                    <h2>${title}</h2>
                    <div class="post-text "><pre class="card-description">${content}</pre></div>
                   <button class="myBtn">Read more</button>
                    <div class="categorie-post">${NameCategories}</div>
                    <div class="content-poster-like">
                        <div class="content-poster">
                            <div class="poster">                                  
                                <img src="assets/user-profile-svgrepo-com.svg" alt="">
                                <div>
                                    <p><span>${Username}</span><img  src="assets/status-active-svgrepo-com.svg" alt=""> </p>
                                    <p>3 weeks ago</p>
                                </div>
                            </div>
                            <div class="like-comment-block">
                                <div>${postDislikCount} dislikes</div>
                                <div>${postLikeCount} Likes</div>
                                <div class="comment">${commentCount} comments <span class="Id-post" style="display:none">${postId}</span> </div>
                            </div>
                        </div>
                    </div>
                
                </div>
            </div> 
         <div class="create-comment ${lastFormComment}"  style="display: none;">
         <form class="form-comment" method="post">
            <input type="hidden" name ="userId" value="${userId}" />
            <input type="hidden" name ="postId" value="${postId}" />

            <textarea type="text" name="ContentComment" id="Content" placeholder="Comment here..."  ></textarea>  
           <button type="submit"> <img src="assets/send-svgrepo-com.svg" alt=""> </button>
         </form>
         </div>
         <div class="bloc-comment ${lastBlocComment}" style="display: none;" >
             ${_comme}
         </div>
    ` ;
    NameCategories="";
 }

    return posts
}

export{
    postsFilter
}