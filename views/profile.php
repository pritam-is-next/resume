<!DOCTYPE html>
<html lang="en" data-bs-theme="dark">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title><?= $$Title ?></title>

    <!-- Bootstrap 5.3 CSS -->
    <link rel="stylesheet" href="/css/Bootstrap/bootstrap.min.css">
    <link rel="stylesheet" href="/css/Bootstrap-Icons/bootstrap-icons.css">

    <style>
        ::-webkit-scrollbar { width: 8px; }
        ::-webkit-scrollbar-track { background: #212529; }
        ::-webkit-scrollbar-thumb {
            background: #495057;
            border-radius: 4px;
        }

        .profile-label {
            font-size: 0.85rem;
            color: #adb5bd;
        }
    </style>
</head>

<body class="bg-body">

<div class="d-flex">

    <!-- Sidebar (unchanged) -->
    <section id="sidebar"
        class="offcanvas-lg offcanvas-start bg-body border-end border-secondary min-vh-100"
        tabindex="-1" style="width: 280px;">
        <div class="offcanvas-body d-flex flex-column h-100 p-3">
            <a href="/admin" class="d-flex align-items-center mb-3 text-body text-decoration-none">
                <span class="fs-4 fw-bold">
                    <i class="bi bi-person-circle me-2"></i><?= $$Heading ?>
                </span>
            </a>

            <ul class="nav nav-pills flex-column mb-auto"></ul>

            <hr class="border-secondary">

            <div class="dropdown">
                <a href="#" class="d-flex align-items-center text-body text-decoration-none dropdown-toggle"
                   data-bs-toggle="dropdown">
                    <img src="<?= $$User_Details->Avatar ?>"
                         class="rounded-circle me-2" width="32">
                    <strong><?= $$User_Details->FullName ?></strong>
                </a>

                <ul class="dropdown-menu shadow">
                    <li>
                        <a class="dropdown-item" href="/admin/profile/edit">
                            <i class="bi bi-pencil-square me-2"></i>Edit Profile
                        </a>
                    </li>
                    <li>
                        <button class="dropdown-item" id="themeToggle">
                            <i class="bi bi-sun-fill me-2"></i>Light Mode
                        </button>
                    </li>
                    <li><hr class="dropdown-divider"></li>
                    <li>
                        <a class="dropdown-item text-danger" href="/logout">
                            <i class="bi bi-box-arrow-right me-2"></i>Logout
                        </a>
                    </li>
                </ul>
            </div>
        </div>
    </section>

    <!-- Main Content -->
    <div class="flex-grow-1 p-4">

        <div class="card bg-body border-secondary shadow-sm">
            <div class="card-header bg-body-tertiary border-secondary d-flex justify-content-between">
                <h5 class="mb-0">
                    <i class="bi bi-person-lines-fill me-2"></i>Profile Overview
                </h5>
                <a href="/admin/profile/edit" class="btn btn-sm btn-outline-primary">
                    <i class="bi bi-pencil me-1"></i>Edit
                </a>
            </div>

            <div class="card-body">
                <div class="row g-4">

                    <!-- Avatar -->
                    <div class="col-md-3 text-center">
                        <img src="<?= $$User_Details->Avatar ?>"
                             class="rounded-circle mb-3"
                             width="140" height="140">
                    </div>

                    <!-- Details -->
                    <div class="col-md-9">

                        <div class="mb-3">
                            <div class="label">Full Name</div>
                            <div><?= $$User_Details->FullName ?></div>
                        </div>

                        <div class="mb-3">
                            <div class="label">Email</div>
                            <div><?= $$User_Details->Email ?></div>
                        </div>

                        <div class="mb-3">
                            <div class="label">Phone</div>
                            <div><?= $$User_Details->Phone  ?></div>
                        </div>

                        <div class="mb-3">
                            <div class="label">Date of Birth</div>
                            <div><?= $$User_Details->Dob  ?></div>
                        </div>

                        <div class="mb-3">
                            <div class="label">Gender</div>
                            <div><?= $$User_Details->Gender  ?></div>
                        </div>

                        <div class="mb-4">
                            <div class="label">Bio</div>
                            <div class="text-muted">
                                <?= $$User_Details->Bio ?>
                            </div>
                        </div>

                        <hr class="border-secondary">

                        <div class="mb-2">
                            <div class="label">Address</div>
                            <div><?= $$User_Details->AddressLine  ?></div>
                        </div>

                        <div class="mb-2">
                            <?= $$User_Details->City ?>,
                            <?= $$User_Details->State ?>
                        </div>

                        <div>
                            <?= $$User_Details->Country ?> - <?= $$User_Details->ZipCode ?>
                        </div>

                    </div>
                </div>
            </div>
        </div>

    </div>
</div>

<!-- Bootstrap JS -->
<script src="/js/Bootstrap/bootstrap.bundle.min.js"></script>

<script>
document.getElementById('themeToggle').addEventListener('click', function () {
    const html = document.documentElement;
    const theme = html.getAttribute('data-bs-theme') === 'dark' ? 'light' : 'dark';
    html.setAttribute('data-bs-theme', theme);
    this.innerHTML = theme === 'dark'
        ? '<i class="bi bi-sun-fill me-2"></i>Light Mode'
        : '<i class="bi bi-moon-stars-fill me-2"></i>Dark Mode';
});
</script>

</body>
</html>
