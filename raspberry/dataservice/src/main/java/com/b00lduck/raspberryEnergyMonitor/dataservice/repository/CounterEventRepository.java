package com.b00lduck.raspberryEnergyMonitor.dataservice.repository;

import com.b00lduck.raspberryEnergyMonitor.dataservice.entity.Counter;
import com.b00lduck.raspberryEnergyMonitor.dataservice.entity.CounterEvent;
import org.springframework.data.repository.CrudRepository;

/**
 * @author Daniel Zerlett (daniel@zerlett.eu)
 * created 24.07.2015
 */
public interface CounterEventRepository extends CrudRepository<CounterEvent, Long> {

}
