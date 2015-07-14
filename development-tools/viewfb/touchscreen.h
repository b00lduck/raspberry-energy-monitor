#pragma once

void ts_uninit();
int ts_init(char* device);
void ts_press(unsigned int x, unsigned y);
void ts_release();

