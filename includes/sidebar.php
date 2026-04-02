<?php
// On récupère le nom de la page actuelle
$currentPage = basename($_SERVER['PHP_SELF']);

// Configuration des liens : 'NOM' => 'FICHIER'
$menuItems = [
    'Dashboard'    => 'admin_dashboard.php',
    'Logistique'   => 'admin_logistique.php',
    'Validation'   => 'admin_validation.php',
    'Finance'      => 'admin_finance.php',
    'Marketing'    => 'admin_marketing.php',
    'Utilisateurs' => 'admin_utilisateurs.php'
];
?>

<aside class="fixed inset-y-0 left-0 w-72 bg-[#1E4E4E] text-white p-8 hidden lg:block shadow-2xl">
    <div class="mb-12">
        <h1 class="font-montserrat font-bold text-xl tracking-tighter">
            <span class="text-[#6A9A9A]">UCU</span> Admin
        </h1>
        <p class="text-[8px] text-[#6A9A9A] font-bold uppercase tracking-[0.3em] mt-1">Gestion Générale</p>
    </div>

    <nav class="space-y-4">
        <p class="text-[9px] font-bold text-white/30 uppercase tracking-[0.2em] mb-6">Menu Principal</p>

        <?php foreach ($menuItems as $label => $file): 
            // Vérification si la page est active pour le style visuel
            $isActive = ($currentPage == $file) 
                ? 'border-r-4 border-[#6A9A9A] bg-white/5 opacity-100' 
                : 'opacity-60 hover:opacity-100 hover:bg-white/5';
        ?>
            <a href="<?= $file ?>" class="block py-3 px-4 text-[10px] font-bold uppercase tracking-[0.2em] transition-all <?= $isActive ?>">
                <?= $label ?>
            </a>
        <?php endforeach; ?>
    </nav>

    <div class="absolute bottom-8 left-8 right-8">
        <div class="pt-6 border-t border-white/10 flex flex-col gap-4">
            <p class="text-[8px] text-white/20 uppercase font-bold italic tracking-widest">
                UpcycleConnect © 2026
            </p>
            <a href="../login/logout.php" class="text-[9px] text-red-400/70 hover:text-red-400 font-bold uppercase tracking-tighter transition-colors">
                Déconnexion
            </a>
        </div>
    </div>
</aside>