% ----------------------------------
% Ejercicio 1: sumList(Lista, Suma).

sumList(L,S):-
    sumList_aux(L,S,0).

sumList_aux([],S,S).

sumList_aux([H|T],S,C):-
    C1 is C+H,
    sumList_aux(T,S,C1).

% ----------------------------------
%  Ejercicio 2:
%--largoPar(Lista).
%--largoImpar(Lista).

largoPar(L):-
    largoPar_aux(L,0).

largoPar_aux([],C):-
    0 is mod(C,2).

largoPar_aux([_|T],C):-
    C1 is C+1,
    largoPar_aux(T,C1).

largoImpar(L):-
    largoImpar_aux(L,0).

largoImpar_aux([],C):-
    not(0 is mod(C,2)).

largoImpar_aux([_|T],C):-
    C1 is C+1,
    largoImpar_aux(T,C1).

% ----------------------------------
% Ejercicio 3: Subconjunto(Lista, Subconjunto).


subConjunto(L, S) :-
   subConjunto_aux(L, S).

subConjunto_aux(L, L).

subConjunto_aux([_|T], S) :-
   subConjunto(T, S).
subConjunto_aux([H|T], [H|S]) :-
   subConjunto_aux(T, S).

% ----------------------------------
% Ejercicio 4: invertir(Lista, Invertida).


invertir(L,I):-
    invertir_aux(L,I,[]).

invertir_aux([],I,I).

invertir_aux([H|T],I,L2):-
    invertir_aux(T,I,[H|L2]).

% ----------------------------------
