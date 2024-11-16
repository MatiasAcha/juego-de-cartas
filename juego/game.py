import random

class Game():
    def __init__(self, tablero, player1, player2):
        self.tablero = tablero
        self.turno = 1
        self.ganador = None
        self.player1 = player1
        self.player2 = player2

        
    def start_game(self):
        def cara_o_seca(self, eleccion_jugador):
            # Verifica que la elección del jugador sea válida
            if eleccion_jugador not in ["cara", "seca"]:
                print("Elección no válida. Debes escoger 'cara' o 'seca'.")
                return None
            
            # Genera un resultado aleatorio entre 1-'cara' o 2-'seca'
            resultado = random.choice(["1", "2"])
            
            # Verifica si el jugador adivinó correctamente
            if eleccion_jugador == resultado:
                return True
            else:
                return False
        print("Elije una opción:")
        print("1. Cara")
        print("2. Seca")
        resultado = cara_o_seca(input())
        if resultado:
            print("¡Has acertado!")
            print("1. Comenzar Primero")
            print("2. Comenzar Segundo")
            res = input()
            if res == "1":    
                for i in range(5):
                    self.player2.robar()
                for i in range(4):
                    self.player1.robar()
                self.player1.iniciarturno()
            if res == "2":
                for i in range(5):
                    self.player1.robar()
                for i in range(4):
                    self.player2.robar()
                self.player2.iniciarturno()
        


    