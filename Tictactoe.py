import os
import random 
import sys
from PyQt5.QtGui import *
from PyQt5.QtWidgets import *
from PyQt5.QtCore import *
os.system("cls")

class Tictac(QMainWindow):
    __ls_btn = []
    __player1 = 0
    __player2 = 0
    __popitka = 1
    __turn = False
    _turn_on = 1
    def __init__(self):
        super().__init__()
        
        self.setFixedSize(800,600)
        self.setStyleSheet("""background-color: #00a8b5;""")
        self.setWindowTitle("Tictac game")
        self.setFont(QFont("Consolas",30))

        self.btn1 = self.make_button("",self)
        self.btn1.move(50,50)

        self.btn2 = self.make_button("",self)
        self.btn2.move(150,50)

        self.btn3 = self.make_button("",self)
        self.btn3.move(250,50)

        self.btn4 = self.make_button("",self)
        self.btn4.move(50,150)

        self.btn5 = self.make_button("",self)
        self.btn5.move(150,150)

        self.btn6 = self.make_button("",self)
        self.btn6.move(250,150)

        self.btn7 = self.make_button("",self)
        self.btn7.move(50,250)

        self.btn8 = self.make_button("",self)
        self.btn8.move(150,250)

        self.btn9 = self.make_button("",self)
        self.btn9.move(250,250)

        self.new_game = QPushButton("New Game",self)
        self.new_game.setGeometry(450,250,200,60)
        self.new_game.setFont(QFont("Consolas",20))
        self.new_game.setStyleSheet("background-color: #de4383")
        self.new_game.clicked.connect(self.start_game)

        self.computer = QPushButton("Play with robot",self)
        self.computer.setGeometry(450,320,200,60)
        self.computer.setFont(QFont("Consolas",14))
        self.computer.setStyleSheet("background-color: #ffe26f")
        self.computer.clicked.connect(self.comp_on)

        self.__ls_btn = [self.btn1,self.btn2,self.btn3,self.btn4,
                         self.btn5,self.btn6,self.btn7,self.btn8,
                         self.btn9]
        
        self.function = [self.click_btn1,self.click_btn2,self.click_btn3,self.click_btn4,
                         self.click_btn5,self.click_btn6,self.click_btn7,self.click_btn8,
                         self.click_btn9]
        
        for x,y in zip(self.__ls_btn,self.function):
            x.clicked.connect(y)

        
        self.playerx = QLabel("Player1:",self)
        self.playerx.setGeometry(400,50,150,50)
        self.playerx.setFont(QFont("Consolas",20))

        self.playero = QLabel("Player2:",self)
        self.playero.setGeometry(400,120,150,50)
        self.playero.setFont(QFont("Consolas",20))
        
        self.tablo = QSpinBox(self)
        self.tablo.setGeometry(600,50,100,50)
        self.tablo.setFont(QFont("Consolas",20))
        self.tablo.setValue(0)
        self.tablo.setEnabled(False)

        self.tablo2 = QSpinBox(self)
        self.tablo2.setGeometry(600,120,100,50)
        self.tablo2.setFont(QFont("Consolas",20))
        self.tablo2.setValue(0)
        self.tablo2.setEnabled(False)
    
    def comp_on(self):
        Tictac._turn_on += 1
        if Tictac._turn_on % 2 == 0:
            Tictac.__turn = True
            Tictac.__popitka = 1
            self.playero.setText("Robot🤖:")
        else:
            Tictac.__turn = False
            self.playero.setText("Player2:")


    def comp_turn(self):
        Tictac.__popitka += 1
        if Tictac.__turn == False:
            return
        comp_btn = [btn for btn in self.__ls_btn if btn.isEnabled()]
        if comp_btn:
            comp_choice = random.choice(comp_btn)
            comp_choice.setText(self.turn())
            self.check_win()
            comp_choice.setEnabled(False)

    def start_game(self):
        Tictac.__player1 = 0
        Tictac.__player2 = 0
        for x in self.__ls_btn:
            x.setText("")
            x.setEnabled(True)
        self.tablo.setValue(Tictac.__player1)
        self.tablo2.setValue(Tictac.__player2)
        
    
    
    def check_win(self):
        wino = False
        winx = False
        if Tictac.__popitka % 2 == 1:
            if self.btn1.text() == 'O' and self.btn2.text() == 'O' and self.btn3.text() == 'O':
                wino = True
            elif self.btn1.text() == 'O' and self.btn5.text() == 'O' and self.btn9.text() == 'O':
                wino = True
            elif self.btn3.text() == 'O' and self.btn5.text() == 'O' and self.btn7.text() == 'O':
                wino = True
            elif self.btn1.text() == 'O' and self.btn4.text() == 'O' and self.btn7.text() == 'O':
                wino = True
            elif self.btn3.text() == 'O' and self.btn6.text() == 'O' and self.btn9.text() == 'O':
                wino = True
            elif self.btn4.text() == 'O' and self.btn5.text() == 'O' and self.btn6.text() == 'O':
                wino = True
            elif self.btn7.text() == 'O' and self.btn8.text() == 'O' and self.btn9.text() == 'O':
                wino = True
            elif self.btn2.text() == 'O' and self.btn5.text() == 'O' and self.btn8.text() == 'O':
                wino = True
        else:
            if self.btn1.text() == 'X' and self.btn2.text() == 'X' and self.btn3.text() == 'X':
                winx = True
            elif self.btn1.text() == 'X' and self.btn5.text() == 'X' and self.btn9.text() == 'X':
                winx = True
            elif self.btn3.text() == 'X' and self.btn5.text() == 'X' and self.btn7.text() == 'X':
                winx = True
            elif self.btn1.text() == 'X' and self.btn4.text() == 'X' and self.btn7.text() == 'X':
                winx = True
            elif self.btn3.text() == 'X' and self.btn6.text() == 'X' and self.btn9.text() == 'X':
                winx = True
            elif self.btn4.text() == 'X' and self.btn5.text() == 'X' and self.btn6.text() == 'X':
                winx = True
            elif self.btn7.text() == 'X' and self.btn8.text() == 'X' and self.btn9.text() == 'X':
                winx = True
            elif self.btn2.text() == 'X' and self.btn5.text() == 'X' and self.btn8.text() == 'X':
                winx = True
        
        if wino:
            msg = QMessageBox(self)
            msg.setText(f"Player2 wins!\nDo you want another round?")
            msg.setIcon(QMessageBox.Information)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("Result!")
            msg.setStandardButtons(QMessageBox.Yes|QMessageBox.No)
            msg.buttonClicked.connect(self.message)
            msg.show()
            Tictac.__player2 += 1
            self.tablo2.setValue(Tictac.__player2)
            for x in self.__ls_btn:
                x.setEnabled(False)
        elif winx:
            msg = QMessageBox(self)
            msg.setText(f"Player1 wins!\nDo you want another round?")
            msg.setIcon(QMessageBox.Information)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("Result!")
            msg.setStandardButtons(QMessageBox.Yes|QMessageBox.No)
            msg.buttonClicked.connect(self.message)
            msg.show()
            Tictac.__player1 += 1
            self.tablo.setValue(Tictac.__player1)
            for x in self.__ls_btn:
                x.setEnabled(False)
        elif Tictac.__popitka == 10 and winx == False and wino == False:
            msg = QMessageBox(self)
            msg.setText(f"Friendship wins!\nDo you want another round?")
            msg.setIcon(QMessageBox.Information)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("Result!")
            msg.setStandardButtons(QMessageBox.Yes|QMessageBox.No)
            msg.buttonClicked.connect(self.message)
            msg.show()
    def message(self,x):
        if x.text() == "&Yes":
            self.start_game()
        else:
            exit()

    def make_button(self,x,y):
        btn = QPushButton(x,y)
        btn.resize(100,100)
        btn.setFont(QFont("Colibri",30))
        btn.setStyleSheet("""border: 2px solid rgb(153, 81, 255);
                             color: rgb(105, 0, 255);
                             border-radius: 15px;
                             background-color: rgb(255, 215, 0);""")
        return btn
    
    def turn(self):
        if Tictac.__popitka % 2 == 0:
            return "X"
        else:
            return "O"

    def click_btn1(self):
        Tictac.__popitka += 1
        if self.btn1.text() == '':
            self.btn1.setText(self.turn())
            self.check_win()
            self.btn1.setEnabled(False)
           
        if Tictac.__turn:
            self.comp_turn()
    
    def click_btn2(self):
        Tictac.__popitka += 1
        if self.btn2.text() == '':
            self.btn2.setText(self.turn())
            self.check_win()
            self.btn2.setEnabled(False)
        if Tictac.__turn:
            self.comp_turn()
    def click_btn3(self):
        Tictac.__popitka += 1
        if self.btn3.text() == '':
            self.btn3.setText(self.turn())
            self.check_win()
            self.btn3.setEnabled(False)
        if Tictac.__turn:
            self.comp_turn()
    def click_btn4(self):
        Tictac.__popitka += 1
        if self.btn4.text() == '':
            self.btn4.setText(self.turn())
            self.check_win()
            self.btn4.setEnabled(False)
        if Tictac.__turn:
            self.comp_turn()
    def click_btn5(self):
        Tictac.__popitka += 1
        if self.btn5.text() == '':
            self.btn5.setText(self.turn())
            self.check_win()
            self.btn5.setEnabled(False)
        if Tictac.__turn:
            self.comp_turn()
    def click_btn6(self):
        Tictac.__popitka += 1
        if self.btn6.text() == '':
            self.btn6.setText(self.turn())
            self.check_win()
            self.btn6.setEnabled(False)
        if Tictac.__turn:
            self.comp_turn()
    def click_btn7(self):
        Tictac.__popitka += 1
        if self.btn7.text() == '':
            self.btn7.setText(self.turn())
            self.check_win()
            self.btn7.setEnabled(False)
        if Tictac.__turn:
            self.comp_turn()
    def click_btn8(self):
        Tictac.__popitka += 1
        if self.btn8.text() == '':
            self.btn8.setText(self.turn())
            self.check_win()
            self.btn8.setEnabled(False)
        if Tictac.__turn:
            self.comp_turn()
    def click_btn9(self):
        Tictac.__popitka += 1
        if self.btn9.text() == '':
            self.btn9.setText(self.turn())
            self.check_win()
            self.btn9.setEnabled(False)
        if Tictac.__turn:
            self.comp_turn()
app = QApplication(sys.argv)
project = Tictac()
project.show()
sys.exit(app.exec_())