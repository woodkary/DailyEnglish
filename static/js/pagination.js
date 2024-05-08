/*
 * @Date: 2024-04-17 14:17:35
 */
let BUTTON_NUM = 5; // 每页显示的按钮数
let page = 1; // 当前页数
let totalPage = 10; // 总页数
//分页按钮
createPaginationButtons();
function createPaginationButtons() {
    const startBtns = document.querySelector(".startBtns");
    let startBtn = document.createElement("button");
    startBtn.classList.add('button');
    let i = document.createElement("i");
    i.classList.add("fa", "fa-angles-left");
    startBtn.appendChild(i);
    startBtn.addEventListener("click", () => {
        page = 1;
        updatePagination(BUTTON_NUM);
    });
    startBtns.appendChild(startBtn);
    let prevButton = document.createElement("button");
    prevButton.classList.add('button');
    let i2 = document.createElement("i");
    i2.classList.add("fa", "fa-angle-left");
    prevButton.appendChild(i2);
    prevButton.addEventListener("click", () => {
        page--;
        if (page < 1) {
            page = 1;
        }
        updatePagination(BUTTON_NUM);
    });
    startBtns.appendChild(prevButton);
    updatePagination(BUTTON_NUM);
    const endBtns = document.querySelector(".endBtns");
    const nextButton = document.createElement("button");
    nextButton.classList.add('button');
    let i3 = document.createElement("i");
    i3.classList.add("fa", "fa-angle-right");
    nextButton.appendChild(i3);
    nextButton.addEventListener("click", () => {
        page++;
        if (page > totalPage) {
            page = totalPage;
        }
        updatePagination(BUTTON_NUM);
    });
    endBtns.appendChild(nextButton);

    const endBtn = document.createElement("button");
    endBtn.classList.add('button');
    let i4 = document.createElement("i");
    i4.classList.add("fa", "fa-angles-right");
    endBtn.appendChild(i4);
    endBtn.addEventListener("click", () => {
        page = totalPage;
        updatePagination(BUTTON_NUM);
    });
    endBtns.appendChild(endBtn);
}

function updatePagination(pageSize) {
    const totalPages = totalPage; // 总页数
    let startPage = Math.max(1, page - Math.floor(pageSize / 2));
    let endPage = Math.min(totalPages, startPage + pageSize - 1);

    if (endPage - startPage < pageSize - 1) {
        startPage = Math.max(1, endPage - pageSize + 1);
    }
    const pagination = document.getElementById('pagination');
    pagination.innerHTML = '';

    for (let i = startPage; i <= endPage; i++) {
        const a = document.createElement('a');
        a.textContent = i;
        a.classList.add('link');
        a.addEventListener('click', () => {
            page = i;
            updatePagination(BUTTON_NUM);
        });

        if (i === page) {
            a.classList.add('active');
        }

        pagination.appendChild(a);
    }
}