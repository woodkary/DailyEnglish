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
    toggleDarkModeSidebar();
    toggleDarkModeNavbar();
    toggleAllSvgs();
    toggleAllImages();
    toggleDarkModeCard2();
    /*toggleAllHtml();*/
    isDarkMode =!isDarkMode;
}
function toggleDarkModeSidebar() {
    var body = document.body;
    var sidebar = document.querySelector('.sidebar');
    var sidebarLinks = document.querySelectorAll('.sidebar a, .dropdown-btn');
    var activeLinks = document.querySelectorAll('.sidebar .active');
    var onThisPageLinks = document.querySelectorAll('.sidebar .onThisPage');
    var dropdownBtn = document.querySelectorAll('.sidebar .dropdown-btn');
    var border = document.querySelector('.sidebar');

    // Check if the body has the 'dark-mode' class
    if (isDarkMode) {
        // Light mode
        body.classList.remove('dark-mode');
        sidebar.style.backgroundColor = '#ffffff';
        sidebarLinks.forEach(function(link) {
            link.style.color = '#818181';
        });
        activeLinks.forEach(function(link) {
            link.style.color = '#818181';
        });
        onThisPageLinks.forEach(function(link) {
            link.style.color = '#6200EE';
        });
        dropdownBtn.forEach(function(btn) {
            btn.style.color = '#818181';
        });
        border.style.borderColor = '#ccc';
    } else {
        // Dark mode
        body.classList.add('dark-mode');
        sidebar.style.backgroundColor = '#121212';
        sidebarLinks.forEach(function(link) {
            link.style.color = '#e0e0e0';
        });
        activeLinks.forEach(function(link) {
            link.style.color = '#e0e0e0';
        });
        onThisPageLinks.forEach(function(link) {
            link.style.color = '#9c27b0';
        });
        dropdownBtn.forEach(function(btn) {
            btn.style.color = '#e0e0e0';
        });
        border.style.borderColor = '#333';
    }
}
function toggleDarkModeNavbar() {
    var navbar = document.querySelector('.navbar');
    var navbarText = document.querySelectorAll('.navbar, .navbar a');
    var navbarBorder = document.querySelector('.navbar');
    var notifications = document.querySelector('.notifications');
    var nightMode = document.querySelector('.night-mode');
    var userInfo = document.querySelector('.user-info');
    var userAvatar = document.querySelector('.user-avatar');
    var logout = document.querySelector('.logout');

    // Check if the body has the 'dark-mode' class
    if (isDarkMode) {
        // Light mode
        navbar.style.backgroundColor = 'white';
        navbarText.forEach(function(element) {
            element.style.color = '#1f2328';
        });
        navbarBorder.style.borderBottom = '1px solid #d8d8d8';
        notifications.onmouseover = function() {
            this.style.backgroundColor = '#d8d8d8';
        };
        notifications.onmouseout = function() {
            this.style.backgroundColor = 'transparent';
        };
        nightMode.style.backgroundColor = 'transparent';
        userInfo.style.fontFamily = "'黑体', serif";
        userInfo.style.color = '#1f2328'; // Assuming the text color needs to change
        userAvatar.style.backgroundColor = 'transparent'; // Assuming the avatar background needs to change
        logout.style.backgroundColor = 'aliceblue';
    } else {
        // Dark mode
        navbar.style.backgroundColor = '#121212';
        navbarText.forEach(function(element) {
            element.style.color = '#e0e0e0';
        });
        navbarBorder.style.borderBottom = '1px solid #333';
        notifications.onmouseover = function() {
            this.style.backgroundColor = '#333';
        };
        notifications.onmouseout = function() {
            this.style.backgroundColor = 'transparent';
        };
        nightMode.style.backgroundColor = 'transparent';
        userInfo.style.fontFamily = "'黑体', serif";
        userInfo.style.color = '#e0e0e0'; // Assuming the text color needs to change
        userAvatar.style.backgroundColor = '#333'; // Assuming the avatar background needs to change
        logout.style.backgroundColor = '#333';
    }
}
function toggleAllSvgs() {
    var svgs = document.querySelectorAll('svg');
    svgs.forEach(function(svg) {
        if(!isDarkMode){
            svg.style.fill = '#e0e0e0';

        }else{
            svg.style.fill = '#1f2328';
        }
    });
}
function toggleAllImages() {
    var images = document.querySelectorAll('img');
    images.forEach(function (image) {
        if(image.classList.contains('not-change')){
            return;
        }
        if (!isDarkMode) {
            image.classList.add('dark');
        } else {
            image.classList.remove('dark');
        }
    });
}

function toggleDarkModeCard2() {
    var card2 = document.querySelectorAll('.card2');
    var card2HoverShadow = document.querySelectorAll('.card2:hover');
    var card2Th = document.querySelectorAll('.card2 th');
    var card2Td = document.querySelectorAll('.card2 td, .card2 th');
    var mainLabel = document.querySelectorAll('.main_label');
    var container = document.querySelectorAll('.container');
    var details1 = document.querySelectorAll('.details_1');
    var title = document.querySelectorAll('.title');
    var count = document.querySelectorAll('.count');
    var notification = document.querySelectorAll('.notification');
    var notificationItem = document.querySelectorAll('.notification-item');
    var notificationItemAfter = document.querySelectorAll('.notification-item:after');
    var notificationItemDivider = document.querySelectorAll('.notification-item-divider');

    // Check if the body has the 'dark-mode' class
    if (isDarkMode) {
        // Light mode
        card2.forEach(function(card) {
            if(card.classList.contains('not-change'))
                return;
            card.style.borderColor = '#ddd';
            card.style.boxShadow = '0 2px 4px rgba(0, 0, 0, 0.1), 4px 0 4px rgba(0, 0, 0, 0.1)';
            card.style.backgroundColor = '#fff';
        });
        card2HoverShadow.forEach(function(card) {
            if(card.classList.contains('not-change'))
                return;
            card.style.boxShadow = '0 4px 8px rgba(0, 0, 0, 0.2)';
        });
        card2Th.forEach(function(th) {
            if(th.classList.contains('not-change'))
                return;
            th.style.backgroundColor = '#f2f2f2';
        });
        card2Td.forEach(function(td) {
            if(td.classList.contains('not-change'))
                return;
            td.style.color = 'black';
        });
        mainLabel.forEach(function(label) {
            if(label.classList.contains('not-change'))
                return;
            label.style.backgroundColor = '#ffffff';
            label.style.color = '#000000';
            label.style.borderLeftColor = '#409eff';
        });
        container.forEach(function(cont) {
            if (cont.classList.contains('not-change'))
                return;
            cont.style.background = 'linear-gradient(to right, #ff7e5f, #feb47b)';
            cont.style.after = 'content: \'\'; position: absolute; left: 0; bottom: 0; width: 100%; height: 15%; background: linear-gradient(to bottom, rgba(255, 255, 255, 0.3), rgba(255, 255, 255, 0.3));';
        });
        details1.forEach(function(det) {
            if(det.classList.contains('not-change'))
                return;
            det.style.color = '#666';
        });
        title.forEach(function(tit) {
            if(tit.classList.contains('not-change'))
                return;
            tit.style.color = '#ffffff';
        });
        count.forEach(function(cou) {
            if(cou.classList.contains('not-change'))
                return;
            cou.style.color = 'white';
        });
        notification.forEach(function(not) {
            if(not.classList.contains('not-change'))
                return;
            not.style.borderColor = '#ccc';
        });
        notificationItem.forEach(function(item) {
            if(item.classList.contains('not-change'))
                return;
            item.style.color = 'black';
        });
        notificationItemAfter.forEach(function(item) {
            if(item.classList.contains('not-change'))
                return;
            item.style.color = '#aeacac';
        });
        notificationItemDivider.forEach(function(div) {
            if(div.classList.contains('not-change'))
                return;
            div.style.borderTopColor = '#ccc';
        });
    } else {
        // Dark mode
        card2.forEach(function(card) {
            if(card.classList.contains('not-change'))
                return;
            card.style.borderColor = '#333';
            card.style.boxShadow = '0 2px 4px rgba(255, 255, 255, 0.1), 4px 0 4px rgba(255, 255, 255, 0.1)';
            card.style.backgroundColor = '#121212';
        });
        card2HoverShadow.forEach(function(card) {
            if(card.classList.contains('not-change'))
                return;
            card.style.boxShadow = '0 4px 8px rgba(255, 255, 255, 0.2)';
        });
        card2Th.forEach(function(th) {
            if(th.classList.contains('not-change'))
                return;
            th.style.backgroundColor = '#333';
        });
        card2Td.forEach(function(td) {
            if(td.classList.contains('not-change'))
                return;
            td.style.color = '#e0e0e0';
        });
        mainLabel.forEach(function(label) {
            if(label.classList.contains('not-change'))
                return;
            label.style.backgroundColor = '#121212';
            label.style.color = '#e0e0e0';
            label.style.borderLeftColor = '#409eff';
        });
        container.forEach(function(cont) {
            if (cont.classList.contains('not-change'))
                return;
            cont.style.background = 'linear-gradient(to right, #3a3a3a, #595959)';
            cont.style.after = 'content: \'\'; position: absolute; left: 0; bottom: 0; width: 100%; height: 15%; background: linear-gradient(to bottom, rgba(0, 0, 0, 0.3), rgba(0, 0, 0, 0.3));';
        });
        details1.forEach(function(det) {
            if(det.classList.contains('not-change'))
                return;
            det.style.color = '#e0e0e0';
        });
        title.forEach(function(tit) {
            if(tit.classList.contains('not-change'))
                return;
            tit.style.color = '#e0e0e0';
        });
        count.forEach(function(cou) {
            if(cou.classList.contains('not-change'))
                return;
            cou.style.color = '#e0e0e0';
        });
        notification.forEach(function(not) {
            if(not.classList.contains('not-change'))
                return;
            not.style.borderColor = '#333';
        });
        notificationItem.forEach(function(item) {
            if(item.classList.contains('not-change'))
                return;
            item.style.color = '#e0e0e0';
        });
        notificationItemAfter.forEach(function(item) {
            if(item.classList.contains('not-change'))
                return;
            item.style.color = '#595959';
        });
        notificationItemDivider.forEach(function(div) {
            if(div.classList.contains('not-change'))
                return;
            div.style.borderTopColor = '#333';
        });
    }
}



