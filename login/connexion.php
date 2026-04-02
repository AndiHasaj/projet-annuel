<?php include '../includes/header.php'; ?>

<div class="min-h-screen flex flex-col items-center justify-center p-6 bg-[#F8FAFC]">
    
    <a href="../accueil.php" class="mb-8 flex flex-col items-center group">
        <span class="font-montserrat font-bold text-2xl text-forest tracking-tighter group-hover:text-sage transition">UPCYCLECONNECT</span>
        <span class="text-[9px] font-bold text-gray-400 uppercase tracking-[0.3em] mt-1">Retour au site</span>
    </a>

    <div class="bg-white w-full max-w-md p-10 rounded-[2.5rem] shadow-2xl border border-gray-100 relative overflow-hidden">
        <div class="absolute top-0 right-0 w-32 h-32 bg-sage/5 rounded-full -mr-16 -mt-16"></div>

        <div class="text-center mb-10">
            <h1 class="font-montserrat text-xl text-forest font-bold uppercase tracking-tight">Espace Personnel</h1>
            <p class="text-gray-400 text-[10px] font-medium mt-2 uppercase tracking-widest">Veuillez vous identifier</p>
        </div>

        <div class="flex bg-gray-50 p-1 rounded-2xl mb-8">
            <button class="flex-1 py-2 text-[9px] font-bold uppercase rounded-xl bg-white shadow-sm text-forest border border-gray-100">Client / Pro</button>
            <button class="flex-1 py-2 text-[9px] font-bold uppercase text-gray-400 hover:text-forest transition">Salarié / Admin</button>
        </div>

        <form action="../traitement_connexion.php" method="POST" class="space-y-5">
            <div>
                <label class="block text-[10px] font-bold text-gray-400 uppercase mb-2 ml-1 tracking-widest">Email ou Identifiant</label>
                <input type="text" name="identifiant" required 
                    placeholder="ex: jean.pierre@mail.com"
                    class="w-full px-6 py-4 rounded-2xl bg-gray-50 border-transparent focus:bg-white focus:ring-2 focus:ring-sage focus:border-transparent outline-none transition text-sm text-forest font-medium">
            </div>

            <div>
                <div class="flex justify-between mb-2 ml-1">
                    <label class="text-[10px] font-bold text-gray-400 uppercase tracking-widest">Mot de passe</label>
                    <a href="#" class="text-[9px] font-bold text-sage hover:underline uppercase">Oublié ?</a>
                </div>
                <input type="password" name="password" required 
                    placeholder="••••••••"
                    class="w-full px-6 py-4 rounded-2xl bg-gray-50 border-transparent focus:bg-white focus:ring-2 focus:ring-sage focus:border-transparent outline-none transition text-sm">
            </div>

            <button type="submit" 
                class="w-full bg-forest text-white font-bold py-5 rounded-2xl hover:bg-sage transition-all transform active:scale-95 shadow-xl shadow-forest/10 uppercase text-[10px] tracking-[0.2em] mt-4">
                Se connecter
            </button>
        </form>

        <div class="mt-10 pt-8 border-t border-gray-50 text-center">
            <p class="text-xs text-gray-500 font-medium">
                Nouveau sur la plateforme ? 
                <a href="inscription.php" class="text-sage font-bold hover:underline ml-1 uppercase text-[10px]">Créer un compte</a>
            </p>
        </div>
    </div>

    <div class="mt-8 flex gap-4 opacity-40">
        <button class="text-[9px] font-bold text-forest uppercase tracking-widest hover:opacity-100 transition">Français</button>
        <button class="text-[9px] font-bold text-forest uppercase tracking-widest hover:opacity-100 transition">English</button>
    </div>
</div>

<?php include '../includes/footer.php'; ?>