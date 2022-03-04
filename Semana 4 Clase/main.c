#include <stdio.h>
#include <malloc.h>

int const SIZE_ARRAY = 10;


void print_array(int *arreglo, int size){
    int *end_ptr = arreglo+size;
    while (arreglo < end_ptr){
        printf("%i ", *arreglo);
        arreglo++;
    }
    printf("\n");
}

void copy_array(int *source_ptr, int *dest_ptr, int size){
    int *end_ptr = source_ptr+size;

    while (source_ptr < end_ptr){
        *dest_ptr = *source_ptr;
        source_ptr++;
        dest_ptr++;
    }
}

int insert_array(int* source_ptr, int ele, int size, int pos){
    source_ptr = realloc(source_ptr,(size+1)*sizeof(int));
    //*(source_ptr+size) = ele;
    // implementar ciclo para insertar en una posición específica que contemple mover todos los elementos posteriores... DEBE INCLUIR EL PARÁMETRO pos
    int *end_pos = source_ptr+size;
    int *desire_pos = &source_ptr[pos];
    while (end_pos != desire_pos){
        *end_pos = *(end_pos-1);
        end_pos--;
    }
    *end_pos = ele;
    return size+1;
}

int main() {
    int arreglo[SIZE_ARRAY];
    int *arreglo2 = (int*)malloc(SIZE_ARRAY*sizeof(int));

    arreglo[0] = 9; arreglo[1] = 8; arreglo[2] = 7; arreglo[3] = 17; arreglo[4] = 99;
    arreglo[5] = 57; arreglo[6] = 25; arreglo[7] = 75; arreglo[8] = 49; arreglo[9] = 36;

    printf("Arreglo Original: ");
    print_array(arreglo, SIZE_ARRAY);

    copy_array(arreglo, arreglo2, SIZE_ARRAY);
    printf("Arreglo Copia: ");
    print_array(arreglo2, SIZE_ARRAY);

    int new_size = 0;
    new_size =  insert_array(arreglo2, 777777, SIZE_ARRAY, 5);
    printf("Arreglo con nuevo elemento en posicion deseada: ");
    print_array(arreglo2, new_size);

    return 0;
}
