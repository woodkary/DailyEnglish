let isDarkMode = false;

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
    /*toggleDarkModeSidebar();
    toggleDarkModeNavbar();
    toggleAllSvgs();
    toggleAllImages();
    toggleDarkModeCard2();
    /!*toggleAllHtml();*!/
    isDarkMode =!isDarkMode;*/
    var options = {
        bottom: '64px', // default: '32px'
        right: 'unset', // default: '32px'
        left: '32px', // default: 'unset'
        time: '0.3s', // default: '0.3s'
        mixColor: '#fff', // default: '#fff'
        backgroundColor: '#fff',  // default: '#fff'
        buttonColorDark: '#100f2c',  // default: '#100f2c'
        buttonColorLight: '#fff', // default: '#fff'
        saveInCookies: false, // default: true,
        label: '', // default: ''
        autoMatchOsTheme: true // default: true
    }
    new Darkmode(options).toggle();
}