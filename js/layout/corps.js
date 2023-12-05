
const leftBloc=(_leftBloc,CatgoryArray)=>{
       _leftBloc.innerHTML=`<div> <img src="assets/category-list-solid-svgrepo-com.svg" alt="">Categories</div>
       `;
       if (CatgoryArray) {
          CatgoryArray.forEach(category => {
            let _div=  document.createElement('div');
            _div.textContent=category.NameCategory;
            _leftBloc.appendChild(_div);
          });
       }
}

const middleBloc=(_middleBloc,_posts)=>{
   
    _middleBloc.appendChild(createPostMenue());

    _middleBloc.appendChild(contentPostBloc(_posts));
    // _middleBloc.appendChild(createComment());


}
const rigthtBloc=(_rigthtBloc)=>{
    _rigthtBloc.innerHTML=`<h2> <img src="assets/right-arrow-svgrepo-com.svg" alt=""><span>Users On line</span> </h2>
                                <div class="bloc">
                                <div class="user">  
                                    <div class="user-infos">                               
                                        <img src="assets/user-profile-svgrepo-com.svg" alt="">
                                        <div>
                                            <p><span>mass</span><img  src="assets/status-active-svgrepo-com.svg" alt=""> </p>                            
                                        </div>
                                    </div> 
                                    <div class="chat-text">
                                        <span>chat</span>
                                        <img src="assets/chat-dots-svgrepo-com.svg" alt="">
                                    </div>
                                </div>
                                <div class="user">  
                                    <div class="user-infos">                               
                                        <img src="assets/user-profile-svgrepo-com.svg" alt="">
                                        <div>
                                            <p><span>abdou</span><img  src="assets/status-active-svgrepo-com.svg" alt=""> </p>                            
                                        </div>
                                    </div> 
                                    <div class="chat-text">
                                        <span>chat</span>
                                        <img src="assets/chat-dots-svgrepo-com.svg" alt="">
                                    </div>
                                </div>
                                <div class="user">  
                                    <div class="user-infos">                               
                                        <img src="assets/user-profile-svgrepo-com.svg" alt="">
                                        <div>
                                            <p><span>adama</span><img  src="assets/status-active-svgrepo-com.svg" alt=""> </p>                            
                                        </div>
                                    </div> 
                                    <div class="chat-text">
                                        <span>chat</span>
                                        <img src="assets/chat-dots-svgrepo-com.svg" alt="">
                                    </div>
                                </div>
                                <div class="user">  
                                    <div class="user-infos">                               
                                        <img src="assets/user-profile-svgrepo-com.svg" alt="">
                                        <div>
                                            <p><span>yassine</span><img  src="assets/status-active-svgrepo-com.svg" alt=""> </p>                            
                                        </div>
                                    </div> 
                                    <div class="chat-text">
                                        <span>chat</span>
                                        <img src="assets/chat-dots-svgrepo-com.svg" alt="">
                                    </div>
                                </div>
                                
                            </div>    
    `;
}

function createPostMenue() {

    let _div=  document.createElement('div');
    _div.className="create-post-menue";
    _div.innerHTML=`<span>
                        <img src="assets/postcard-svgrepo-com.svg" alt="">
                        <span>My Posts</span>
                    </span>
                    <button class="create-post-btn">Create Post</button>
    `;
    return _div;
}

function createComment() {
    let _div=  document.createElement('div');
    _div.className="create-comment";
    // _div.innerHTML=`<span>
    //                     <img src="assets/postcard-svgrepo-com.svg" alt="">
    //                     <span>My Posts</span>
    //                 </span>
    //                 <button class="create-post-btn">Create Post</button>
    // `;
    return _div;
}

function contentPostBloc(_posts) {
  
  let _div=  document.createElement('div');
    _div.className="content-post-block";
    _div.innerHTML=posts(_posts);
    return _div;
}

function posts(_posts) {
    var posts="";
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
        const postId=post["Posts"]["ID"];
console.log(postId);
        // recupere es cqtory du post
        let NameCategories="";
        for (let j = 0; j < Categories.length; j++) {
            const NameCategory = Categories[j]["NameCategory"];
            NameCategories+=`<span>${NameCategory}</span>`;
        }
        
        // console.log(`post ${i} : ${post}`);
        let image=Image ? `<div class="image-post"><img src="/static/image/${Image}" alt=""></div>` : "";
        
     posts += `<div class="one-post-block">
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
         <div class="create-comment">
         <form>
            <textarea type="text" name="Content" id="Content" placeholder="Comment here..."  ></textarea>  
           <button type="submite"> <img src="assets/send-svgrepo-com.svg" alt=""> </button>
         </form>
         </div>
    ` ;
    NameCategories="";
 }

    return posts
}


export {leftBloc,middleBloc,rigthtBloc}