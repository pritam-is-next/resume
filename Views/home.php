<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>
        <?= $Title ?> | Portfolio
    </title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="/Css/Bootstrap/bootstrap.min.css">
    <link rel="stylesheet" href="/Css/Bootstrap-Icons/bootstrap-icons.css">
    <link rel="stylesheet" href="/Css/main.css">
</head>

<body>
    <header>
        <nav class="navbar navbar-expand-lg bg-body-tertiary">
            <div class="container-fluid">
                <a class="navbar-brand" href="/">
                    <?= $Heading ?>
                </a>
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
                    data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent"
                    aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul class="navbar-nav ms-auto mb-2 mb-lg-0">
                        <?php foreach ($NavItems as $key => $NavItem): ?> <!-- Nav Items Range -->
                            <?php if ($NavItem->DropDown): ?>
                                <li class="nav-item dropdown">
                                    <a class="nav-link dropdown-toggle" href="<?= $NavItem->Href ?>" role="button"
                                        data-bs-toggle="dropdown" aria-expanded="false"><?= $NavItem->Name ?></a>
                                    <ul class="dropdown-menu">
                                        <?php foreach ($DropDown as $key => $DropDownItem): ?>
                                            <li><a class="dropdown-item" href="<?= $DropDownItem->Href ?>"><?= $DropDownItem->Name ?>
                                                </a></li>
                                            <?php if (count($NavItem->DropDown)): ?>
                                                break <?= count($NavItem->DropDown) ?>
                                            <?php endif; ?>
                                        <?php endforeach ?>

                                    </ul>
                                </li>
                            <?php elseif ($NavItem->Disabled): ?>
                                <li class="nav-item">
                                    <a class="nav-link disabled" aria-disabled="true"><?= $NavItem->Name ?> </a>
                                </li>
                            <?php else: ?>
                                <li class="nav-item">
                                    <a class="nav-link active" aria-current="page"
                                        href="<?= $NavItem->Href ?> "><?= $NavItem->Name ?> </a>
                                </li>
                            <?php endif ?> <!-- If End-->
                        <?php endforeach ?> <!-- Nav Items Range -->
                    </ul>
                </div>

            </div>
        </nav>
    </header>
    <main>
        <div class="col-sm-3">

        </div>
    </main>
    <!-- Footer -->
    <footer class="footer"> </footer>

    <!-- Bootstrap JS -->
    <script src="/Js/Bootstrap/bootstrap.bundle.min.js"></script>
</body>

</html>