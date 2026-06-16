<?php include '../includes/header.php'; ?>

<div class="min-h-screen flex flex-col items-center justify-center p-6 bg-[#F8FAFC] py-12">
    
    <div class="text-center mb-10">
        <h1 class="font-montserrat text-3xl text-forest font-bold uppercase tracking-tighter">Nous rejoindre</h1>
        <p class="text-gray-400 text-[10px] font-bold mt-2 uppercase tracking-[0.2em]">Choisissez votre profil UpcycleConnect</p>
    </div>

    <div class="bg-white w-full max-w-2xl p-10 rounded-[3rem] shadow-2xl border border-gray-100">
        
        <form action="../register/traitement_inscription_particulier.php" method="POST" class="space-y-8">

            <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
                <div class="space-y-2">
                    <label class="text-[9px] font-bold text-gray-400 uppercase tracking-widest ml-1">Nom</label>
                    <input type="text" name="nom" required class="w-full px-6 py-4 rounded-2xl bg-gray-50 border-transparent focus:bg-white focus:ring-2 focus:ring-sage outline-none transition text-sm">
                </div>
                <div class="space-y-2">
                    <label class="text-[9px] font-bold text-gray-400 uppercase tracking-widest ml-1">Prénom</label>
                    <input type="text" name="prenom" required class="w-full px-6 py-4 rounded-2xl bg-gray-50 border-transparent focus:bg-white focus:ring-2 focus:ring-sage outline-none transition text-sm">
                </div>
                <div class="space-y-2">
                    <label class="text-[9px] font-bold text-gray-400 uppercase tracking-widest ml-1">Email</label>
                    <input type="email" name="email" required class="w-full px-6 py-4 rounded-2xl bg-gray-50 border-transparent focus:bg-white focus:ring-2 focus:ring-sage outline-none text-sm">
                </div>
                <div class="space-y-2">
                    <label class="text-[9px] font-bold text-gray-400 uppercase tracking-widest ml-1">Mot de passe</label>
                    <input type="password" name="password" required class="w-full px-6 py-4 rounded-2xl bg-gray-50 border-transparent focus:bg-white focus:ring-2 focus:ring-sage outline-none text-sm">
                </div>
            </div>

            <div class="pt-4">
                <button type="submit" class="w-full bg-forest text-white font-bold py-5 rounded-2xl hover:bg-sage transition-all transform active:scale-95 shadow-xl uppercase text-[10px] tracking-[0.2em]">
                    Créer mon compte
                </button>
                <p class="text-center mt-6 text-[10px] text-gray-400 font-bold uppercase tracking-widest">
                    Déjà inscrit ? <a href="../login/connexion.php" class="text-sage hover:underline">Se connecter</a>
                </p>
            </div>
        </form>
    </div>
</div>

<?php include '../includes/footer.php'; ?>