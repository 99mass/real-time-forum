
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

const middleBloc=(_middleBloc)=>{

    _middleBloc.appendChild(createPostMenue())

    _middleBloc.appendChild(contentPostBloc())

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

function contentPostBloc() {

  let _div=  document.createElement('div');
    _div.className="content-post-block";
    _div.innerHTML=posts();
    return _div;
}

function posts() {
    return `<div class="one-post-block">
                <div class="image-post">
                    <img src="assets/img/im.jpeg" alt="">
                </div>
                <div class="post-content">
                    <h2>Blockchain developer best practices on innovationchain</h2>
                    <div class="post-text card-description">
                        Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, sint est mollitia fugit cupiditate corrupti doloribus, 
                        laboriosam ad libero beatae numquam quasi saepe recusandae maxime earum, expedita eligendi eos perferendis.Lorem ipsum dolor sit, 
                        amet consectetur adipisicing elit. Quo laborum, reprehenderit expedita atque quia officiis natus voluptatum inventore unde dolor ad laboriosam illo cupiditate, ipsum, nam provident vero corporis magni!logout-svgrepo-com
                        Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, sint est mollitia fugit cupiditate corrupti doloribus, 
                        
                        </div>
                   <button class="myBtn">Read more</button>
                    <div class="categorie-post"><span>Sport</span><span>Education</span><span>Politic</span></div>
                    <div class="content-poster-like">
                        <div class="content-poster">
                            <div class="poster">                                  
                                <img src="assets/user-profile-svgrepo-com.svg" alt="">
                                <div>
                                    <p><span>osamb</span><img  src="assets/status-active-svgrepo-com.svg" alt=""> </p>
                                    <p>3 weeks ago</p>
                                </div>
                            </div>
                            <div class="like-comment-block">
                                <div>651,324 dislikes</div>
                                <div>36,6545 Likes</div>
                                <div>56 comments</div>
                            </div>
                        </div>
                    </div>
                
                </div>
            </div> 
            
  

    `;
}


export {leftBloc,middleBloc,rigthtBloc}