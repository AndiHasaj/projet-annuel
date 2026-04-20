<?php include '../includes/header.php'; ?>

<div class="min-h-screen flex flex-col items-center justify-center p-6 bg-[#F8FAFC] py-12">
    
    <div class="text-center mb-10">
        <h1 class="font-montserrat text-3xl text-forest font-bold uppercase tracking-tighter">Nous rejoindre</h1>
        <p class="text-gray-400 text-[10px] font-bold mt-2 uppercase tracking-[0.2em]">Choisissez votre profil UpcycleConnect</p>
    </div>

    <div class="bg-white w-full max-w-2xl p-10 rounded-[3rem] shadow-2xl border border-gray-100">
        
        <form id="inscriptionForm" method="POST" class="space-y-8">
            <script>
                document.getElementById("inscriptionForm").addEventListener("submit", function(e) {

                    const role = document.querySelector('input[name="role"]:checked').value;

                    if (role === "particulier") {
                        this.action = "inscription_particulier.php";
                    } 
                    else if (role === "pro") {
                        this.action = "inscription_professionnel.php";
                    }
                });
            </script>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <label class="relative cursor-pointer group">
                    <input type="radio" name="role" value="particulier" class="peer hidden" checked>
                    <div class="p-6 border-2 rounded-[2rem] transition-all peer-checked:border-sage peer-checked:bg-sage/5 group-hover:border-sage/20 bg-gray-50/50">
                        <div class="flex justify-between items-start mb-4">
                            <span class="text-3xl">🌱</span>
                            <div class="w-5 h-5 rounded-full border-2 border-gray-300 peer-checked:border-sage peer-checked:bg-sage"></div>
                        </div>
                        <p class="font-montserrat font-bold text-forest text-sm uppercase">Particulier</p>
                        <p class="text-[10px] text-gray-400 mt-2 leading-relaxed">Je souhaite donner des objets, participer à des formations et suivre mon impact.</p>
                    </div>
                </label>

                <label class="relative cursor-pointer group">
                    <input type="radio" name="role" value="pro" class="peer hidden">
                    <div class="p-6 border-2 rounded-[2rem] transition-all peer-checked:border-sage peer-checked:bg-sage/5 group-hover:border-sage/20 bg-gray-50/50">
                        <div class="flex justify-between items-start mb-4">
                            <span class="text-3xl">🛠️</span>
                            <div class="w-5 h-5 rounded-full border-2 border-gray-300"></div>
                        </div>
                        <p class="font-montserrat font-bold text-forest text-sm uppercase">Professionnel</p>
                        <p class="text-[10px] text-gray-400 mt-2 leading-relaxed">Artisan ou Entreprise : je souhaite récupérer des matériaux et gérer mes contrats.</p>
                    </div>
                </label>
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