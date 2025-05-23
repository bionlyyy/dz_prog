f = open('input.txt', encoding="utf-8").read()
key = 3
def caesar(text, shift): #цезарь
    result = ''
    for s in text:
        if 'a' <= s <= 'z':  #англ мелкие
            result += chr(((ord(s) - ord('a') + shift) % 26) + ord('a'))
        elif 'A' <= s <= 'Z':  #англ большие
            result += chr(((ord(s) - ord('A') + shift) % 26) + ord('A'))
        elif 'а' <= s <= 'я':  #рус мeлкие
            result += chr(((ord(s) - ord('а') + shift) % 32) + ord('а'))
        elif 'А' <= s <= 'Я':  #рус большие
            result += chr(((ord(s) - ord('А') + shift) % 32) + ord('А'))
        else:
            result += s
    return result

def atbash(text): #атбаш
    result = ''
    for s in text:
        if 'a' <= s <= 'z': #англ больш и мелк
            result += chr(ord('a') + ord('z') - ord(s))
        elif 'A' <= s <= 'Z':
            result += chr(ord('A') + ord('Z') - ord(s))
        elif "а" <= s <= "я": #рус больш и мелк
            result += chr(ord("а") + ord("я") - ord(s))
        elif "А" <= s <= "Я":
            result += chr(ord("А") + ord("Я") - ord(s))
        else:
            result += s
    return result
shif = caesar(f, key)
deshif = caesar(shif, -key)
sshif = atbash(f)
desshif = atbash(sshif)
with open('output.txt', 'w', encoding='utf-8') as file:
    file.write(f"Исходный текст:\n{f}\n\n")
    file.write(f"Зашифровано Цезарем :\n{shif}\n\n")
    file.write(f"Расшифровано Цезарем:\n{deshif}\n\n")
    file.write(f"Зашифровано Атбашем:\n{sshif}\n\n")
    file.write(f"Расшифровано Атбашем:\n{desshif}\n")