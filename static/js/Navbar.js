function toggleMoon() {
    let moon = document.getElementById("moon");
    let startIndex=moon.src.lastIndexOf("/")+1;
    let endIndex=moon.src.lastIndexOf(".");
    let moonUrl=moon.src.slice(startIndex,endIndex);
    if(moonUrl==="moon") {
        moon.src = "./image/sun.svg";
    }else{
        moon.src = "./image/moon.svg";
    }
}