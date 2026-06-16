<?php include 'includes/header.php'; ?>
<div class="flex">
    <nav class="w-20 bg-[#1E4E4E] h-screen flex flex-col items-center py-10 gap-8">
        <div class="w-10 h-10 bg-[#6A9A9A] rounded-xl flex items-center justify-center text-white">UC</div>
        <div class="text-white/40 hover:text-white cursor-pointer">📦</div>
        <div class="text-white/40 hover:text-white cursor-pointer">📄</div>
        <div class="text-white/40 hover:text-white cursor-pointer">🔔</div>
    </nav>

    <main class="flex-1 p-12 bg-[#F8FAFC]">
        <h2 class="font-montserrat text-3xl text-[#1E4E4E] mb-10">Espace <span class="text-[#6A9A9A]">Artisan</span></h2>
        
        <div class="grid grid-cols-2 gap-10">
            <div class="bg-white p-10 rounded-[2.5rem] shadow-sm">
                <h3 class="font-bold uppercase text-xs tracking-widest text-gray-400 mb-6">Mon Abonnement Pro</h3>
                <div class="flex items-center justify-between border-b pb-4 mb-4">
                    <p class="text-sm font-bold text-[#1E4E4E]">Abonnement Premium</p>
                    <span class="text-[#6A9A9A] font-bold">ACTIF</span>
                </div>
                <button class="w-full py-4 bg-gray-50 text-[10px] font-bold uppercase rounded-xl">Télécharger Facture PDF</button>
            </div>

            <div class="bg-white p-10 rounded-[2.5rem] shadow-sm">
                <h3 class="font-bold uppercase text-xs tracking-widest text-gray-400 mb-6">Objets réservés</h3>
                <div class="p-4 bg-orange-50 border-l-4 border-orange-400 rounded-xl">
                    <p class="text-xs font-bold text-orange-700">Lot de bois de chêne - Box #104</p>
                    <p class="text-[10px] text-orange-600 mt-1">Code de récupération : UCU-9921</p>
                </div>
            </div>
        </div>
    </main>
</div>