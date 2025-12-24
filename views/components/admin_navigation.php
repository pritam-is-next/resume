<hr class="border-secondary">
<ul class="nav nav-pills flex-column mb-auto">
    <li class="nav-item">
        <?php foreach ($$Nav_items as $key => $nav_item): ?>
            <?php if ($key == 0): ?>

                <a href="<?= $nav_item->Href ?>" class="nav-link active bg-body-secondary text-body" aria-current="page">
                    <?= $nav_item->Name ?>
                </a>

            <?php else: ?>

                <a href="<?= $nav_item->Href ?>" class="nav-link text-body">
                    <?= $nav_item->Name ?>
                </a>

            <?php endif ?>

        <?php endforeach ?>


    </li>
</ul>