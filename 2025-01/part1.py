class Dial:
    def __init__(self):
        self.dial: int = 50
        self.temp: int = 0
        self.zeroCount: int = 0
        self.currentOp: str = ""
        
    def moveAmount(self) -> int:
        direction: str = self.currentOp[0]
        amt: int = int(self.currentOp[1:])

        if amt % 100 == 0:
            return 0
        while amt > 100:
            amt -= 100
        while amt < -99:
            amt += 100
        if direction == 'L':
            return -amt
        return amt 
    
    def resolvePosition(self) -> int:  
        if self.temp == 0:
            self.zeroCount += 1
            return self.temp
        elif self.temp < 0:
            return self.temp + 100
        elif self.temp > 99:
            return self.temp - 100
        else:
            return self.temp
        
    def move(self, op: str) -> None:    
        self.currentOp = op 
        self.temp = self.dial + self.moveAmount()
        self.dial = self.resolvePosition()

def main():
    dial = Dial()
    with open("input.txt", "r") as file:
        operations = file.read().split('\n') 

    for op in operations:
        dial.move(op)
    
    print(dial.zeroCount)

if __name__ == "__main__":
    main()