class Tablero:
    def __init__(self):
        # Inicializamos un tablero de 3x3 con valores None
        self.campo = [[None for _ in range(3)] for _ in range(3)]