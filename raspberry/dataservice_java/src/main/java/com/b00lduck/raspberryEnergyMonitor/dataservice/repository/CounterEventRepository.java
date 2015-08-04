package com.b00lduck.raspberryEnergyMonitor.dataservice.repository;

import com.b00lduck.raspberryEnergyMonitor.dataservice.entity.Counter;
import com.b00lduck.raspberryEnergyMonitor.dataservice.entity.CounterEvent;
import org.joda.time.DateTime;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;

import java.util.List;

/**
 * @author Daniel Zerlett (daniel@zerlett.eu)
 * created 24.07.2015
 */
public interface CounterEventRepository extends CrudRepository<CounterEvent, Long> {

	//@Query("select c from CounterEvent")
	//List<CounterEvent> findTimeRange(DateTime startDate, DateTime endDate);

}
