#include <mpi.h>
#include <stdio.h>

#define NUM_ROWS 12
#define NUM_COLS 12

void transferHeat(int time, int* temp, int* room);
void copyRoomTemps(int* room, int* temp);
void printRoom(int* room);
int getTempCur(int* room, int row, int col);
int getTempTop(int* room, int row, int col);
int getTempBot(int* room, int row, int col);
int getTempLeft(int* room, int row, int col);
int getTempRight(int* room, int row, int col);
int decFloor(int floor, int num);
int incCeiling(int ceiling, int num);

int main(int argc, char** argv) {
    // Initialize Room Matrix
    int temp[NUM_ROWS * NUM_COLS] = {0};
    int room[NUM_ROWS * NUM_COLS] = {0};
    temp[3 * NUM_COLS + 3] = 80;
    temp[10 * NUM_COLS + 10] = 80;
    room[3 * NUM_COLS + 3] = 80;
    room[10 * NUM_COLS + 10] = 80;

    // Start Heat transfer for 20 'seconds'
    for (int time = 0; time < 20; time++) {
        transferHeat(time, temp, room);
        copyRoomTemps(room, temp);
    }

    printRoom(room);

    return 0;
}

void transferHeat(int time, int* temp, int* room) {
    for (int i = 0; i < NUM_ROWS; i++) {
        for (int j = 0; j < NUM_COLS; j++) {
            temp[i * NUM_COLS + j] = getTempCur(room, i, j) + 
                                        (getTempTop(room, i, j) +
                                        getTempBot(room, i, j) + 
                                        getTempLeft(room, i, j)+
                                        getTempRight(room, i, j) - 
                                        4*room[i*NUM_COLS+j]) / 4;
        }
    }
}

int getTempCur(int* room, int row, int col) {
    return room[row * NUM_COLS + col];
}

int getTempTop(int* room, int row, int col) {
    int topIndex = decFloor(0, row);
    return room[topIndex * NUM_COLS + col];
}
int getTempBot(int* room, int row, int col) {
    int botIndex = incCeiling(NUM_ROWS, row);
    return room[botIndex * NUM_COLS + col];
}

int getTempLeft(int* room, int row, int col) {
    int leftIndex = decFloor(0, col);
    return room[row * NUM_COLS + leftIndex];
}

int getTempRight(int* room, int row, int col) {
    int rightIndex = incCeiling(NUM_COLS, col);
    return room[row * NUM_COLS + rightIndex];
}

int decFloor(int floor, int num) {
    if (num - 1 < floor) {
        return floor;
    }
    return num - 1;
}

int incCeiling(int ceiling, int num) {
    if (num + 1 >= ceiling) {
        return ceiling - 1;
    }
    return num + 1;
}

void copyRoomTemps(int* room, int* temp) {
    for (int i = 0; i < NUM_ROWS * NUM_COLS; i++) {
        room[i] = temp[i];
    }
}

void printRoom(int* room) {
    for (int i = 0; i < NUM_ROWS; i++) {
        for (int j = 0; j < NUM_COLS; j++) {
            printf("%2d  ", room[i * NUM_COLS + j]);
        }
        printf("\n");
    }
    printf("\n----------------------------------------------------------------\n\n");
}