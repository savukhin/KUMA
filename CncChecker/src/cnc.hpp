#pragma once

#include <stdint.h>

enum class CncStatus {
    Stopped=1,
    Working=2,
    Broken=3
};

int statusToInt(CncStatus status) {
    if (status == CncStatus::Stopped) return 1;
    if (status == CncStatus::Working) return 2;
    return 3;
}

class Cnc {
private:
    uint8_t pin_stopped_;
    uint8_t pin_working_;
    uint8_t pin_broken_;

    bool has_broken_;

    void (*onChange_) (CncStatus&);

    CncStatus status_;

public:
    Cnc(uint8_t pin_stopped, uint8_t pin_working, uint8_t pin_broken): 
        pin_stopped_(pin_stopped), pin_working_(pin_working), pin_broken_(pin_broken), has_broken_(true) {}

    Cnc(uint8_t pin_stopped=1, uint8_t pin_working=2): 
        pin_stopped_(pin_stopped), pin_working_(pin_working), pin_broken_(-1), has_broken_(false) {}

    void addOnChange(void (*callback) (CncStatus&)) {
        this->onChange_ = callback;
    }

    void setStatus(CncStatus new_status) {
        this->status_ = new_status;
        if (this->onChange_) {
            this->onChange_(new_status);
        }
    }

    void loop() {
        bool stopped = digitalRead(pin_stopped_) == HIGH;
        bool working = digitalRead(pin_stopped_) == HIGH;
        bool broken = digitalRead(pin_stopped_) == HIGH;

        if (broken && this->status_ != CncStatus::Broken) {
            setStatus(CncStatus::Broken);
        } else if (stopped && this->status_ != CncStatus::Stopped) {
            setStatus(CncStatus::Stopped);
        } else if (working && this->status_ != CncStatus::Working) {
            setStatus(CncStatus::Working);
        } 
    }

    CncStatus getStatus() {
        return this->status_;
    }
};
