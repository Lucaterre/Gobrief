import sys
from textwrap import TextWrapper
import textwrap
from collections import defaultdict

import spacy
from spacy.symbols import nsubj, VERB, ADJ


input_ten_words = sys.argv[1]

input_text = sys.argv[2]

top_ten_words = input_ten_words.split(",")


#print(f'\n\n {top_ten_words}')
#print(type(top_ten_words))

#print(input_text_without_stops_words)
#print(type(input_text_without_stops_words))






# On charge un modèle français
# ! préalable télécharger le modèle (réseau convolutionnel entraîné sur deux corpus, WikiNER et Sequoia) !
# python -m spacy download fr_core_news_sm
model_fr = spacy.load("fr_core_news_sm")


def return_POS(text):
    """
    """
    document = model_fr(text)
    lexical_field_verbose = defaultdict(list)
    a= []
    for token in document:
        if token.text in top_ten_words:
            top_word = token.text
            head_word = token.head.text
            type_head_word = token.head.pos_
            relation_top_word_and_head_word = token.dep_
            a.append((top_word, f'{head_word} [{type_head_word}] ({relation_top_word_and_head_word})'))
            #lexical_field_verbose[top_word].append(head_word + ' [' + type_head_word + ']' + ' ' + '(' + relation_top_word_and_head_word + ')')
            #print(f'word : {top_word} associate to vocabulary : {head_word} ({type_head_word}) ({relation_top_word_and_head_word})')

    for k, v in a:
        lexical_field_verbose[k].append(v)

    #print(lexical_field_verbose)
    result = sorted(lexical_field_verbose.items())


    print('\nII-/ Lexical field associate with top ten words : \n')
    print(' Index    Word                   Lexical Field                  ')
    print('======= ========= ==============================================')
    index = 1
    for i in result:
        top = index
        prefix = str(top) + "      | "
        preferredWidth = 70
        wrapper = TextWrapper(initial_indent=prefix,
                              width=preferredWidth,
                              subsequent_indent='   '*len(prefix))
        message = f'{i[0]} |  {str(i[1])}'
        print(f'{wrapper.fill(message)}\n')
        index += 1

def report_association():
    pass


    #topics_noun_adj = {}
    #for i, token in enumerate(document):
    #tok_clean = token.text.lower()
    #if tok_clean in top_ten_words:
    #print(tok_clean)
    #if token.pos_ not in ('NOUN', 'PROPN'):
    #continue
    #for j in range(i+1,len(document)):
    #if document[j].pos_ == 'ADJ':
    #topics_noun_adj[token.text] = document[j]
    #break
    #noun_adj_pairs.append((token,document[j]))


return_POS(input_text)
