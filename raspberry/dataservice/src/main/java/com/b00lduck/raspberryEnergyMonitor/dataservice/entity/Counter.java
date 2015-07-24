package com.b00lduck.raspberryEnergyMonitor.dataservice.entity;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;

/**
 * @author Daniel Zerlett (daniel@zerlett.eu)
 * @created 24.07.2015
 */
@Entity
public class Counter {

	@Id
	@GeneratedValue(strategy = GenerationType.TABLE)
	private Long id;

	private String name;

}
