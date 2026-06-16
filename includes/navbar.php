<?php include 'langue_setup.php'; ?>

<nav class="flex justify-between items-center px-10 py-6 border-b sticky top-0 bg-white z-50">
    <div class="font-montserrat font-bold text-2xl text-forest uppercase tracking-tighter">
        UpcycleConnect
    </div>

    <div class="flex items-center space-x-8">
        <div class="space-x-8 text-[11px] font-bold uppercase tracking-widest text-gray-500">
            <a href="/accueil.php" class="hover:text-sage transition"><?php echo $textes['nav_home']; ?></a>
            <a href="/login/connexion.php" class="bg-forest text-white px-6 py-3 rounded-full hover:bg-sage transition">
                <?php echo $textes['nav_login']; ?>
            </a>
        </div>

        <div class="flex gap-2 border-l pl-6">
            <a href="?lang=fr" class="opacity-<?php echo $lang == 'fr' ? '100' : '40' ?> hover:opacity-100">🇫🇷</a>
            <a href="?lang=en" class="opacity-<?php echo $lang == 'en' ? '100' : '40' ?> hover:opacity-100">🇬🇧</a>
        </div>
    </div>
</nav>