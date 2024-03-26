function Navbar(navbarName,dropdownOptions){
    return `<nav class="navbar navbar-expand-lg bg-body-tertiary">
            <div class="container-fluid">
                <a class="navbar-brand">navbarName</a>
                <button class="navbar-toggler" type="button" data-coreui-toggle="collapse"
                        data-coreui-target="#navbarSupportedContent" aria-controls="navbarSupportedContent"
                        aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                        <li class="nav-item">
                            <a class="nav-link active" aria-current="page">首页</a>
                        </li>
                        <li class="nav-item dropdown">
                            <a class="nav-link dropdown-toggle" href="#" role="button" data-coreui-toggle="dropdown"
                               aria-expanded="false">
                                操作
                            </a>
                            <ul class="dropdown-menu">
                                {dropdownOptions.map((option, index) => (
                                    <li key={index} class="dropdown-item">
                                        option
                                    </li>
                                ))}
                            </ul>
                        </li>
                    </ul>
                    <form class="d-flex" role="search">
                        <input class="form-control me-2" type="search" placeholder="Search" aria-label="Search"/>
                        <button class="btn btn-outline-success" type="submit">查找</button>
                    </form>
                </div>
            </div>
        </nav>
    `;
}