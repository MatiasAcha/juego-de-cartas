class Puntos():
    def __init__(self, arr, ab, iz, der):
        self.arriba = arr
        self.abajo = ab
        self.izquierda = iz
        self.derecha = der
    def __str__(self):
        return f"Arriba: {self.arriba}, Abajo: {self.abajo}, Izquierda: {self.izquierda}, Derecha: {self.derecha}"