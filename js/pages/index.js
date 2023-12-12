import { header } from "../layout/header.js";
import { leftBloc, middleBloc, rigthtBloc } from "../layout/corps.js";
import { body, Header, corpsContent, menuGauche, menuMilieu, menuDroite } from "../helper/bigBlocContent.js";

const indexPage = (data) => {
    console.log(data);
  
        var socket = new WebSocket("ws://localhost:8080/ws");
        socket.onopen = () => {
            socket.send(JSON.stringify({
                Username: data["User"]["username"]
            }));
            console.log("socket: "+data["User"]["username"]);
        }
    

    let posts = data["Datas"] ? data["Datas"] : "";

    let CatgoryArray = data["Category"];
    let Username = data["User"]["username"];
    body.appendChild(Header);
    corpsContent.appendChild(menuGauche);
    corpsContent.appendChild(menuMilieu);
    corpsContent.appendChild(menuDroite);
    body.appendChild(corpsContent);

    const _header = document.querySelector(' header');
    const _leftBloc = document.querySelector(' .menu-gauche');
    const _middleBloc = document.querySelector(' .milieu');
    const _rigthtBloc = document.querySelector(' .menu-droite');

    header(_header, Username);
    leftBloc(_leftBloc, CatgoryArray);
    middleBloc(_middleBloc, posts);
    rigthtBloc(_rigthtBloc);

}

export { indexPage }