<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <script src="https://cdn.tailwindcss.com"></script>
<script>
    // Configuration de tes couleurs personnalisées (Mission 1)
    tailwind.config = {
      theme: {
        extend: {
          colors: {
            foret: '#1B4332',
            sauge: '#B7E4C7',
          }
        }
      }
    }
</script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600&family=Montserrat:wght@700&display=swap" rel="stylesheet">
    
    <script>
        tailwind.config = {
            theme: {
                extend: {
                    fontFamily: {
                        sans: ['Inter', 'sans-serif'],
                        montserrat: ['Montserrat', 'sans-serif'],
                    },
                    colors: {
                        forest: '#1E4E4E',
                        sage: '#6A9A9A',
                    }
                }
            }
        }
    </script>
    <title>UpcycleConnect</title>
</head>
<body class="bg-[#F8FAFC] font-sans"></body>